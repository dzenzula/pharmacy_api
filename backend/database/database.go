package database

import (
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPharmacyDB() {
	var err error
	dsn := "sqlserver://sa:123@localhost:1434?database=PHARMACY"
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
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
