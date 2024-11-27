package config

import (
	"log"

	"pharmacy/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var PHDB *gorm.DB

func InitAuthDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("pharmacy_auth_database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	// Автоматическая миграция
	DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.RolePermission{})
}
