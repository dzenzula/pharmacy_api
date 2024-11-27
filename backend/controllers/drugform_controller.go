package controllers

import (
	"net/http"
	"pharmacy/database"
	"pharmacy/models"

	"github.com/gin-gonic/gin"
)

func CreateDrugForm(c *gin.Context) {
	var drugForm models.DrugForm
	if err := c.ShouldBindJSON(&drugForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&drugForm).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create drug form"})
		return
	}
	c.JSON(http.StatusCreated, drugForm)
}

func GetDrugForms(c *gin.Context) {
	var drugForms []models.DrugForm
	if err := database.DB.Find(&drugForms).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch drug forms"})
		return
	}
	c.JSON(http.StatusOK, drugForms)
}

func UpdateDrugForm(c *gin.Context) {
	var drugForm models.DrugForm
	if err := database.DB.First(&drugForm, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Drug form not found"})
		return
	}

	if err := c.ShouldBindJSON(&drugForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&drugForm).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update drug form"})
		return
	}
	c.JSON(http.StatusOK, drugForm)
}

func DeleteDrugForm(c *gin.Context) {
	if err := database.DB.Delete(&models.DrugForm{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete drug form"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Drug form deleted successfully"})
}
