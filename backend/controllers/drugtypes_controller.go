package controllers

import (
	"net/http"
	"pharmacy/database"
	"pharmacy/models"

	"github.com/gin-gonic/gin"
)

func CreateDrugType(c *gin.Context) {
	var drugType models.DrugType
	if err := c.ShouldBindJSON(&drugType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&drugType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create drug type"})
		return
	}
	c.JSON(http.StatusCreated, drugType)
}

func GetDrugTypes(c *gin.Context) {
	var drugTypes []models.DrugType
	if err := database.DB.Find(&drugTypes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch drug types"})
		return
	}
	c.JSON(http.StatusOK, drugTypes)
}

func UpdateDrugType(c *gin.Context) {
	var drugType models.DrugType
	if err := database.DB.First(&drugType, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Drug type not found"})
		return
	}

	if err := c.ShouldBindJSON(&drugType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&drugType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update drug type"})
		return
	}
	c.JSON(http.StatusOK, drugType)
}

func DeleteDrugType(c *gin.Context) {
	if err := database.DB.Delete(&models.DrugType{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete drug type"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Drug type deleted successfully"})
}
