package database

import (
	"log"
	"pharmacy/models"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var AuthDB *gorm.DB

func InitPharmacyDB() {
	var err error
	dsn := "sqlserver://pharmacy_user:pharmacy@localhost:1433?database=PHARMACY&TrustServerCertificate=true&connect_timeout=30"
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
	})
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	// Миграция таблиц
	/*err = phDB.AutoMigrate(
		&models.Manufacturers{},
		&models.DrugType{},
		&models.DrugForm{},
		&models.Drug{},
		&models.Shipment{},
		&models.ShipmentDetail{},
		&models.ShipmentStatus{},
	)
	if err != nil {
		log.Fatalf("Не удалось выполнить миграцию базы данных: %v", err)
	}*/
}

func InitAuthDB() {
	var err error
	dsn := "sqlserver://auth_user:auth@localhost:1433?database=Auth&TrustServerCertificate=true&connect_timeout=30"
	AuthDB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:   true,
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	// Автоматическая миграция
	AuthDB.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.RolePermission{})
}
