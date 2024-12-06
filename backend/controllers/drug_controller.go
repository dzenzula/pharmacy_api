package controllers

import (
	"net/http"
	"pharmacy/database"
	"pharmacy/models"

	"github.com/gin-gonic/gin"
)

func CreateDrug(c *gin.Context) {
	var drug models.Drug
	if err := c.ShouldBindJSON(&drug); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&drug).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create drug"})
		return
	}
	c.JSON(http.StatusCreated, drug)
}

func GetDrugs(c *gin.Context) {
	var drugs []models.Drug
	if err := database.DB.Preload("Type").Preload("Form").Preload("Manufacturer").Find(&drugs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch drugs"})
		return
	}
	c.JSON(http.StatusOK, drugs)
}

func UpdateDrug(c *gin.Context) {
	var drug models.Drug
	if err := database.DB.First(&drug, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Drug not found"})
		return
	}

	if err := c.ShouldBindJSON(&drug); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&drug).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update drug"})
		return
	}
	c.JSON(http.StatusOK, drug)
}

func DeleteDrug(c *gin.Context) {
	if err := database.DB.Delete(&models.Drug{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete drug"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Drug deleted successfully"})
}
