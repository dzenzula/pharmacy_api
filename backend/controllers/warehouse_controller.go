package controllers

import (
	"net/http"
	"pharmacy/database"
	"pharmacy/models"

	"github.com/gin-gonic/gin"
)

// CreateWarehouse - Создать запись на складе
func CreateWarehouse(c *gin.Context) {
	var warehouse models.Warehouse
	if err := c.ShouldBindJSON(&warehouse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentTime := formatCurrentTime()
	warehouse.TimestampInsert = currentTime

	if err := database.DB.Create(&warehouse).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create warehouse entry"})
		return
	}

	c.JSON(http.StatusCreated, warehouse)
}

// GetWarehouses - Получить список записей на складе
func GetWarehouses(c *gin.Context) {
	var warehouses []models.Warehouse
	if err := database.DB.Preload("Drug").Preload("Manufacturer").Find(&warehouses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch warehouses"})
		return
	}

	for i := range warehouses {
		date := warehouses[i].Timestamp
		formatDateString(&date)
	}

	c.JSON(http.StatusOK, warehouses)
}

// UpdateWarehouse - Обновить запись на складе
func UpdateWarehouse(c *gin.Context) {
	id := c.Param("id")
	var warehouse models.Warehouse

	if err := database.DB.First(&warehouse, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Warehouse entry not found"})
		return
	}

	var input models.Warehouse
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var drug models.Drug
	if err := database.DB.First(&drug, input.DrugID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid DrugID"})
		return
	}

	var manufacturer models.Manufacturers
	if err := database.DB.First(&manufacturer, input.ManufacturerID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ManufacturerID"})
		return
	}

	if err := database.DB.Model(&warehouse).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update warehouse entry"})
		return
	}

	c.JSON(http.StatusOK, warehouse)
}

// DeleteWarehouse - Удалить запись на складе
func DeleteWarehouse(c *gin.Context) {
	id := c.Param("id")

	if err := database.DB.Delete(&models.Warehouse{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete warehouse entry"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Warehouse entry deleted successfully"})
}
