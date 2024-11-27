package controllers

import (
	"net/http"
	"pharmacy/database"
	"pharmacy/models"

	"github.com/gin-gonic/gin"
)

func CreateShipment(c *gin.Context) {
	var shipment models.Shipment
	if err := c.ShouldBindJSON(&shipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&shipment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create shipment"})
		return
	}
	c.JSON(http.StatusCreated, shipment)
}

func GetShipments(c *gin.Context) {
	var shipments []models.Shipment
	if err := database.DB.Find(&shipments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch shipments"})
		return
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
