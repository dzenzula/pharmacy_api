package models

import "time"

// Manufacturer - производитель
type Manufacturers struct {
	ManufacturerID    int       `gorm:"column:ManufacturerID;primaryKey"`
	ManufacturerName  string    `gorm:"column:ManufacturerName;type:nvarchar(255);not null"`
	Country           string    `gorm:"column:Country;type:nvarchar(100)"`
	City              string    `gorm:"column:City;type:nvarchar(100)"`
	Address           string    `gorm:"column:Address;type:nvarchar(255)"`
	PhoneNumber       string    `gorm:"column:PhoneNumber;type:nvarchar(50)"`
	Email             string    `gorm:"column:Email;type:nvarchar(100)"`
	Website           string    `gorm:"column:Website;type:nvarchar(255)"`
	LicenseNumber     string    `gorm:"column:LicenseNumber;type:nvarchar(100)"`
	LicenseExpiration time.Time `gorm:"column:LicenseExpiration"`
	AdditionalNotes   string    `gorm:"column:AdditionalNotes;type:text"`
}

func (Manufacturers) TableName() string {
	return "dbo.Manufacturers"
}

// DrugType - тип лекарств
type DrugType struct {
	TypeID      int    `gorm:"column:TypeID;primaryKey"`
	TypeName    string `gorm:"column:TypeName;type:nvarchar(100);not null"`
	Description string `gorm:"column:Description;type:text"`
}

func (DrugType) TableName() string {
	return "dbo.DrugTypes"
}

// DrugForm - форма лекарств
type DrugForm struct {
	FormID      int    `gorm:"column:FormID;primaryKey"`
	FormName    string `gorm:"column:FormName;type:nvarchar(100);not null"`
	Description string `gorm:"column:Description;type:text"`
}

func (DrugForm) TableName() string {
	return "dbo.DrugForms"
}

// Drug - лекарства
type Drug struct {
	DrugID               int       `gorm:"column:DrugID;primaryKey"`
	DrugName             string    `gorm:"column:DrugName;type:nvarchar(255);not null"`
	TypeID               int       `gorm:"column:TypeID;not null"`
	FormID               int       `gorm:"column:FormID;not null"`
	ManufacturerID       int       `gorm:"column:ManufacturerID;not null"`
	ActiveIngredient     string    `gorm:"column:ActiveIngredient;type:nvarchar(255)"`
	Concentration        string    `gorm:"column:Concentration;type:nvarchar(100)"`
	PackSize             string    `gorm:"column:PackSize;type:nvarchar(100)"`
	Dosage               string    `gorm:"column:Dosage;type:nvarchar(100)"`
	PrescriptionRequired bool      `gorm:"column:PrescriptionRequired;default:false"`
	StorageConditions    string    `gorm:"column:StorageConditions;type:nvarchar(255)"`
	ShelfLife            string    `gorm:"column:ShelfLife;type:nvarchar(50)"`
	YearReleased         int       `gorm:"column:YearReleased;not null"`
	TimestampRegistered  time.Time `gorm:"column:TimestampRegistered"`
	TimestampLastUpdate  time.Time `gorm:"column:TimestampLastUpdate"`
	Price                float64   `gorm:"column:Price;type:decimal(20,8)"`
	AdditionalNotes      string    `gorm:"column:AdditionalNotes;type:text"`

	// Связи
	Type         DrugType      `gorm:"foreignKey:TypeID;references:TypeID"`
	Form         DrugForm      `gorm:"foreignKey:FormID;references:FormID"`
	Manufacturer Manufacturers `gorm:"foreignKey:ManufacturerID;references:ManufacturerID"`
}

func (Drug) TableName() string {
	return "dbo.Drugs"
}

// Shipment - поставки
type Shipment struct {
	ID                   int     `gorm:"primaryKey"`
	BatchID              string  `gorm:"column:BatchID;type:nvarchar(50);not null"`
	FromCity             string  `gorm:"column:FromCity;type:nvarchar(100)"`
	FromWarehouseAddress string  `gorm:"column:FromWarehouseAddress;type:nvarchar(255)"`
	ToCity               string  `gorm:"column:ToCity;type:nvarchar(100)"`
	ToWarehouseAddress   string  `gorm:"column:ToWarehouseAddress;type:nvarchar(255)"`
	ToCountry            string  `gorm:"column:ToCountry;type:nvarchar(100)"`
	Timestamp            *string `gorm:"column:Timestamp"`
	LastUpdate           *string `gorm:"column:LastUpdate"`

	// Связи
	Details []ShipmentDetail `gorm:"foreignKey:BatchID;references:BatchID"`
	Status  ShipmentStatus   `gorm:"foreignKey:BatchID;references:BatchID"`
}

func (Shipment) TableName() string {
	return "dbo.Shipments"
}

// ShipmentDetail - детали поставок
type ShipmentDetail struct {
	ID                int     `gorm:"column:ID;primaryKey"`
	BatchID           string  `gorm:"column:BatchID;not null"`
	DrugID            uint    `gorm:"column:DrugID;not null"`
	ManufacturerID    uint    `gorm:"column:ManufacturerID;not null"`
	Quantity          int     `gorm:"column:Quantity;not null"`
	Price             float64 `gorm:"column:Price;type:decimal(20,8)"`
	ManufacturingDate string  `gorm:"column:ManufacturingDate"`

	// Связи
	Drug         Drug          `gorm:"foreignKey:DrugID;references:DrugID"`
	Manufacturer Manufacturers `gorm:"foreignKey:ManufacturerID;references:ManufacturerID"`
}

func (ShipmentDetail) TableName() string {
	return "dbo.ShipmentDetails"
}

// ShipmentStatus - статус поставок
type ShipmentStatus struct {
	ID            int        `gorm:"column:ID;primaryKey"`
	BatchID       string     `gorm:"column:BatchID;not null"`
	DepartureTime *time.Time `gorm:"column:DepartureTime"`
	ArrivalTime   *time.Time `gorm:"column:ArrivalTime"`
	LastUpdate    *string    `gorm:"column:LastUpdate"`
	Status        string     `gorm:"column:Status;type:nvarchar(100)"`
}

func (ShipmentStatus) TableName() string {
	return "dbo.ShipmentStatus"
}

// Warehouse - склад поставок
type Warehouse struct {
	ID              int    `gorm:"column:ID;primaryKey"`
	DrugID          int    `gorm:"column:DrugID;not null"`
	ManufacturerID  int    `gorm:"column:ManufacturerID;not null"`
	BatchNumber     string `gorm:"column:BatchNumber;type:nvarchar(50);not null"`
	Quantity        int    `gorm:"column:Quantity;not null"`
	CertificateCode string `gorm:"column:CertificateCode;type:nvarchar(100)"`
	Timestamp       string `gorm:"column:Timestamp"`
	TimestampInsert string `gorm:"column:TimestampInsert" json:"-"`

	// Связи
	Drug         Drug          `gorm:"foreignKey:DrugID;references:DrugID"`
	Manufacturer Manufacturers `gorm:"foreignKey:ManufacturerID;references:ManufacturerID"`
}

func (Warehouse) TableName() string {
	return "dbo.Warehouse"
}
