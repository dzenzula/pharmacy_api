package config

import (
	"log"

	"pharmacy/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBConfig struct {
	DSN string
}

var dbConfigs = map[string]DBConfig{
	"authConnection": {DSN: "sqlserver://admin:admin@localhost:1434?database=Auth"},
	"pharConnection": {DSN: "sqlserver://admin:admin@localhost:1434?database=PHARMACY"},
}

var DB *gorm.DB

func InitAuthDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("pharmacy_auth_database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	// Автоматическая миграция
	DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.RolePermission{})
}
