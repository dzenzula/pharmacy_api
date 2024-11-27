package models

import "time"

// Manufacturer - производитель
type Manufacturers struct {
	ManufacturerID    uint   `gorm:"primaryKey"`
	ManufacturerName  string `gorm:"type:nvarchar(255);not null;unique"`
	Country           string `gorm:"type:nvarchar(100)"`
	City              string `gorm:"type:nvarchar(100)"`
	Address           string `gorm:"type:nvarchar(255)"`
	PhoneNumber       string `gorm:"type:nvarchar(50)"`
	Email             string `gorm:"type:nvarchar(100);unique"`
	Website           string `gorm:"type:nvarchar(255)"`
	LicenseNumber     string `gorm:"type:nvarchar(100)"`
	LicenseExpiration time.Time
	AdditionalNotes   string `gorm:"type:text"`
}

// DrugType - тип лекарств
type DrugType struct {
	TypeID      uint   `gorm:"primaryKey"`
	TypeName    string `gorm:"type:nvarchar(100);not null;unique"`
	Description string `gorm:"type:text"`
}

// DrugForm - форма лекарств
type DrugForm struct {
	FormID      uint   `gorm:"primaryKey"`
	FormName    string `gorm:"type:nvarchar(100);not null;unique"`
	Description string `gorm:"type:text"`
}

// Drug - лекарства
type Drug struct {
	DrugID               uint   `gorm:"primaryKey"`
	DrugName             string `gorm:"type:nvarchar(255);not null;unique"`
	TypeID               uint   `gorm:"not null"`
	FormID               uint   `gorm:"not null"`
	ManufacturerID       uint   `gorm:"not null"`
	ActiveIngredient     string `gorm:"type:nvarchar(255)"`
	Concentration        string `gorm:"type:nvarchar(100)"`
	PackSize             string `gorm:"type:nvarchar(100)"`
	Dosage               string `gorm:"type:nvarchar(100)"`
	PrescriptionRequired bool   `gorm:"default:false"`
	StorageConditions    string `gorm:"type:nvarchar(255)"`
	ShelfLife            string `gorm:"type:nvarchar(50)"`
	YearReleased         int    `gorm:"not null"`
	TimestampRegistered  time.Time
	TimestampLastUpdate  time.Time
	Price                float64
	AdditionalNotes      string `gorm:"type:text"`

	// Связи
	Type         DrugType      `gorm:"foreignKey:TypeID"`
	Form         DrugForm      `gorm:"foreignKey:FormID"`
	Manufacturer Manufacturers `gorm:"foreignKey:ManufacturerID"`
}

// Shipment - поставки
type Shipment struct {
	ID                   uint   `gorm:"primaryKey"`
	BatchID              uint   `gorm:"not null"`
	FromCity             string `gorm:"type:nvarchar(100)"`
	FromWarehouseAddress string `gorm:"type:nvarchar(255)"`
	ToCity               string `gorm:"type:nvarchar(100)"`
	ToWarehouseAddress   string `gorm:"type:nvarchar(255)"`
	ToCountry            string `gorm:"type:nvarchar(100)"`
	Timestamp            time.Time
	LastUpdate           time.Time

	// Связи
	Details []ShipmentDetail `gorm:"foreignKey:BatchID"`
	Status  []ShipmentStatus `gorm:"foreignKey:BatchID"`
}

// ShipmentDetail - детали поставок
type ShipmentDetail struct {
	ID                uint    `gorm:"primaryKey"`
	BatchID           uint    `gorm:"not null"`
	DrugID            uint    `gorm:"not null"`
	ManufacturerID    uint    `gorm:"not null"`
	Quantity          int     `gorm:"not null"`
	Price             float64 `gorm:"type:decimal(20,8)"`
	ManufacturingDate time.Time

	// Связи
	Drug         Drug          `gorm:"foreignKey:DrugID"`
	Manufacturer Manufacturers `gorm:"foreignKey:ManufacturerID"`
}

// ShipmentStatus - статус поставок
type ShipmentStatus struct {
	ID            uint `gorm:"primaryKey"`
	BatchID       uint `gorm:"not null"`
	DepartureTime time.Time
	ArrivalTime   time.Time
	LastUpdate    time.Time
	Status        string `gorm:"type:nvarchar(100)"`
}
