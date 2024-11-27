package controllers

import (
	"net/http"
	"pharmacy/database"
	"pharmacy/models"

	"github.com/gin-gonic/gin"
)

func CreateShipmentDetail(c *gin.Context) {
	var shipmentDetail models.ShipmentDetail
	if err := c.ShouldBindJSON(&shipmentDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&shipmentDetail).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create shipment detail"})
		return
	}
	c.JSON(http.StatusCreated, shipmentDetail)
}

func GetShipmentDetails(c *gin.Context) {
	var shipmentDetails []models.ShipmentDetail
	if err := database.DB.Find(&shipmentDetails).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch shipment details"})
		return
	}
	c.JSON(http.StatusOK, shipmentDetails)
}

func UpdateShipmentDetail(c *gin.Context) {
	var shipmentDetail models.ShipmentDetail
	if err := database.DB.First(&shipmentDetail, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shipment detail not found"})
		return
	}

	if err := c.ShouldBindJSON(&shipmentDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&shipmentDetail).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update shipment detail"})
		return
	}
	c.JSON(http.StatusOK, shipmentDetail)
}

func DeleteShipmentDetail(c *gin.Context) {
	if err := database.DB.Delete(&models.ShipmentDetail{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete shipment detail"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Shipment detail deleted successfully"})
}
