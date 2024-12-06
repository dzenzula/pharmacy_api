package controllers

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"pharmacy/blockchain"
	"pharmacy/database"
	"pharmacy/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

func formatCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05") // Format as "YYYY-MM-DD HH:mm:ss"
}

func formatDateString(dateStr *string) {
	if dateStr == nil || *dateStr == "" {
		return
	}
	layout := "2006-01-02T15:04:05Z" // ISO 8601 format
	t, err := time.Parse(layout, *dateStr)
	if err != nil {
		return // Keep the original string if parsing fails
	}
	*dateStr = t.Format("2006-01-02 15:04:05") // Format as "YYYY-MM-DD HH:mm:ss"
}

func CreateShipment(c *gin.Context) {
	var shipment models.Shipment
	if err := c.ShouldBindJSON(&shipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	currentTime := formatCurrentTime()
	shipment.LastUpdate = &currentTime
	shipment.Timestamp = &currentTime
	batch := uuid.New().String()
	shipment.BatchID = batch
	shipment.Status = models.ShipmentStatus{
		BatchID: batch,
		Status:  "Accepted",
	}

	hashcode, err := blockchain.AddShipmentToBlockchain(batch, shipment.FromCity, shipment.ToCity, shipment.Status.Status, currentTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add shipment to blockchain"})
		return
	}

	if err := database.DB.Create(&shipment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create shipment"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"shipment": shipment,
		"hashcode": hashcode,
	})
}

func GetShipments(c *gin.Context) {
	var shipments []models.Shipment
	if err := database.DB.Order(clause.OrderByColumn{Column: clause.Column{Name: "ID"}, Desc: true}).Preload("Status").Preload("Details.Drug").Preload("Details.Drug.Manufacturer").Find(&shipments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch shipments"})
		return
	}

	for i := range shipments {
		formatDateString(shipments[i].Timestamp)
		formatDateString(shipments[i].LastUpdate)
	}

	c.JSON(http.StatusOK, shipments)
}

func UpdateShipment(c *gin.Context) {
	var shipment models.Shipment
	if err := database.DB.First(&shipment, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shipment not found"})
		return
	}

	if err := c.ShouldBindJSON(&shipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentTime := formatCurrentTime()

	if shipment.Status.Status != "" {
		shipment.Status.LastUpdate = &currentTime
		if err := database.DB.Save(&shipment.Status).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update shipment status"})
			return
		}
	}

	shipment.LastUpdate = &currentTime
	shipment.Details = []models.ShipmentDetail{}
	if err := database.DB.Save(&shipment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update shipment"})
		return
	}
	c.JSON(http.StatusOK, shipment)
}

func DeleteShipment(c *gin.Context) {
	if err := database.DB.Delete(&models.Shipment{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete shipment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Shipment deleted successfully"})
}

func ProcessQRCode(c *gin.Context) {
	var qrCode struct {
		HashCode string          `json:"hashcode"`
		Shipment models.Shipment `json:"shipment"`
	}

	// Bind JSON data to qrCode struct
	if err := c.ShouldBindJSON(&qrCode); err != nil {
		log.Println("Error binding JSON:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid QR code data"})
		return
	}
	// Validate required shipment fields
	missingFields := []string{}
	if qrCode.Shipment.BatchID == "" {
		missingFields = append(missingFields, "BatchID")
	}
	if qrCode.Shipment.FromCity == "" {
		missingFields = append(missingFields, "FromCity")
	}
	if qrCode.Shipment.ToCity == "" {
		missingFields = append(missingFields, "ToCity")
	}
	if qrCode.Shipment.ToCountry == "" {
		missingFields = append(missingFields, "ToCountry")
	}

	if len(missingFields) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields", "fields": missingFields})
		return
	}

	// Fetch existing shipment from the database
	var existingShipment models.Shipment
	if err := database.DB.Preload("Status").First(&existingShipment, "BatchID = ?", qrCode.Shipment.BatchID).Error; err != nil {
		log.Println("Shipment not found:", err.Error())
		c.JSON(http.StatusNotFound, gin.H{"error": "Shipment not found"})
		return
	}

	// Compare QR data with database data
	mismatchedFields := compareShipments(qrCode.Shipment, existingShipment)
	if len(mismatchedFields) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":             "QR code data does not match database records",
			"mismatched_fields": mismatchedFields,
		})
		return
	}

	// Verify transaction on blockchain
	transactionValid, err := blockchain.VerifyTransaction(qrCode.HashCode)
	if err != nil {
		log.Println("Blockchain verification failed:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify transaction on blockchain"})
		return
	}
	if !transactionValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or unconfirmed blockchain transaction"})
		return
	}

	// Update shipment status on blockchain
	currentTime := formatCurrentTime()
	txHash, err := blockchain.UpdateShipmentOnBlockchain(qrCode.Shipment.BatchID)
	if err != nil {
		log.Println("Blockchain update failed:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update shipment on blockchain"})
		return
	}

	// Synchronize updated data in the local database
	existingShipment.Status.Status = "Delivered"
	*existingShipment.Status.ArrivalTime = time.Now()
	existingShipment.Status.LastUpdate = &currentTime
	existingShipment.LastUpdate = &currentTime

	// Save updated status in the database
	if err := database.DB.Save(&existingShipment.Status).Error; err != nil {
		log.Println("Failed to update shipment status in database:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update shipment status in database"})
		return
	}

	// Save the updated shipment in the database
	if err := database.DB.Save(&existingShipment).Error; err != nil {
		log.Println("Failed to update shipment in database:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update shipment in database"})
		return
	}

	// Return success response with the blockchain transaction hash
	c.JSON(http.StatusOK, gin.H{
		"message":     "Shipment updated successfully",
		"transaction": txHash,
	})
}

// Utility function to check if a string is hexadecimal
func isHexadecimal(s string) bool {
	_, err := hex.DecodeString(s)
	return err == nil
}

// Compare QR code shipment data with database shipment data
func compareShipments(qrShipment, dbShipment models.Shipment) map[string]string {
	mismatchedFields := make(map[string]string)

	if qrShipment.FromCity != dbShipment.FromCity {
		mismatchedFields["FromCity"] = fmt.Sprintf("QR: %s, DB: %s", qrShipment.FromCity, dbShipment.FromCity)
	}
	if qrShipment.ToCity != dbShipment.ToCity {
		mismatchedFields["ToCity"] = fmt.Sprintf("QR: %s, DB: %s", qrShipment.ToCity, dbShipment.ToCity)
	}
	if qrShipment.ToCountry != dbShipment.ToCountry {
		mismatchedFields["ToCountry"] = fmt.Sprintf("QR: %s, DB: %s", qrShipment.ToCountry, dbShipment.ToCountry)
	}
	if qrShipment.FromWarehouseAddress != dbShipment.FromWarehouseAddress {
		mismatchedFields["FromWarehouseAddress"] = fmt.Sprintf("QR: %s, DB: %s", qrShipment.FromWarehouseAddress, dbShipment.FromWarehouseAddress)
	}
	if qrShipment.ToWarehouseAddress != dbShipment.ToWarehouseAddress {
		mismatchedFields["ToWarehouseAddress"] = fmt.Sprintf("QR: %s, DB: %s", qrShipment.ToWarehouseAddress, dbShipment.ToWarehouseAddress)
	}

	return mismatchedFields
}
