package models

import "time"

// User модель
type User struct {
	UserID           uint      `gorm:"column:UserID;primaryKey;autoIncrement"`         // Идентификатор пользователя
	Username         string    `gorm:"column:Username;size:50;unique;not null"`        // Уникальное имя пользователя
	PasswordHash     string    `gorm:"column:PasswordHash;size:255;not null" json:"-"` // Хеш пароля
	Email            string    `gorm:"column:Email;size:100;unique"`                   // Уникальный email
	RoleID           uint      `gorm:"column:RoleID;not null"`                         // Ссылка на роль
	DataAdditionalID uint      `gorm:"column:DataAdditionalID"`
	CreatedAt        time.Time `gorm:"column:CreatedAt;autoCreateTime" json:"-"` // Время создания

	Role           Role           `gorm:"foreignKey:RoleID;references:RoleID"`
	DataAdditional DataAdditional `gorm:"foreignKey:DataAdditionalID;references:DataAdditionalID"`
}

func (User) TableName() string {
	return "dbo.Users"
}

// Role модель
type Role struct {
	RoleID   uint   `gorm:"column:RoleID;primaryKey;autoIncrement"`  // Идентификатор роли
	RoleName string `gorm:"column:RoleName;size:50;unique;not null"` // Уникальное имя роли

	// Связь многие-ко-многим с Permission через таблицу RolePermission
	Permissions []Permission `gorm:"many2many:RolePermission;foreignKey:RoleID;joinForeignKey:RoleID;References:PermissionID;joinReferences:PermissionID"`
}

func (Role) TableName() string {
	return "dbo.Roles"
}

// Permission модель
type Permission struct {
	PermissionID   uint   `gorm:"column:PermissionID;primaryKey;autoIncrement"`   // Идентификатор разрешения
	PermissionName string `gorm:"column:PermissionName;size:100;unique;not null"` // Уникальное имя разрешения
	Description    string `gorm:"column:Description" json:"description"`          // Описание разрешения

	Roles []Role `gorm:"many2many:RolePermission;foreignKey:PermissionID;joinForeignKey:PermissionID;References:RoleID;joinReferences:RoleID" json:"-"`
}

func (Permission) TableName() string {
	return "dbo.Permissions"
}

// RolePermission модель для связи многие-ко-многим
type RolePermission struct {
	RoleID       uint `gorm:"column:RoleID;primaryKey"`       // Идентификатор роли
	PermissionID uint `gorm:"column:PermissionID;primaryKey"` // Идентификатор разрешения
}

func (RolePermission) TableName() string {
	return "dbo.RolePermission"
}

// AdditionalData дополнительная информация о пользователе
type DataAdditional struct {
	DataAdditionalID uint   `gorm:"column:DataAdditionalID;primaryKey;autoIncrement"`
	ManufacturerName string `gorm:"column:ManufacturerName;size:100"`
}

func (DataAdditional) TableName() string {
	return "dbo.DataAdditional"
}
