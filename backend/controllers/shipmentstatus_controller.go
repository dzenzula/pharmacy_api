package controllers

import (
	"net/http"
	"pharmacy/database"
	"pharmacy/models"

	"github.com/gin-gonic/gin"
)

func CreateShipmentStatus(c *gin.Context) {
	var shipmentStatus models.ShipmentStatus
	if err := c.ShouldBindJSON(&shipmentStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&shipmentStatus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create shipment status"})
		return
	}
	c.JSON(http.StatusCreated, shipmentStatus)
}

func GetShipmentStatuses(c *gin.Context) {
	var shipmentStatuses []models.ShipmentStatus
	if err := database.DB.Find(&shipmentStatuses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch shipment statuses"})
		return
	}
	c.JSON(http.StatusOK, shipmentStatuses)
}

func UpdateShipmentStatus(c *gin.Context) {
	var shipmentStatus models.ShipmentStatus
	if err := database.DB.First(&shipmentStatus, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shipment status not found"})
		return
	}

	if err := c.ShouldBindJSON(&shipmentStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&shipmentStatus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update shipment status"})
		return
	}
	c.JSON(http.StatusOK, shipmentStatus)
}

func DeleteShipmentStatus(c *gin.Context) {
	if err := database.DB.Delete(&models.ShipmentStatus{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete shipment status"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Shipment status deleted successfully"})
}
