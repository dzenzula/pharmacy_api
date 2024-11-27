package controllers

import (
	"net/http"
	"pharmacy/database"
	"pharmacy/models"

	"github.com/gin-gonic/gin"
)

func CreateManufacturer(c *gin.Context) {
	var manufacturer models.Manufacturers
	if err := c.ShouldBindJSON(&manufacturer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&manufacturer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create manufacturer"})
		return
	}
	c.JSON(http.StatusCreated, manufacturer)
}

func GetManufacturers(c *gin.Context) {
	var manufacturers []models.Manufacturers
	if err := database.DB.Find(&manufacturers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch manufacturers"})
		return
	}
	c.JSON(http.StatusOK, manufacturers)
}

func UpdateManufacturer(c *gin.Context) {
	var manufacturer models.Manufacturers
	if err := database.DB.First(&manufacturer, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Manufacturer not found"})
		return
	}

	if err := c.ShouldBindJSON(&manufacturer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&manufacturer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update manufacturer"})
		return
	}
	c.JSON(http.StatusOK, manufacturer)
}

func DeleteManufacturer(c *gin.Context) {
	if err := database.DB.Delete(&models.Manufacturers{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete manufacturer"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Manufacturer deleted successfully"})
}
