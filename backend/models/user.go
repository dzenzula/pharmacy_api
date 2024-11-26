package models

import "time"

// User модель
type User struct {
	UserID       uint      `gorm:"column:userid;primaryKey;autoIncrement" json:"userId"`
	Username     string    `gorm:"column:username;size:50;unique;not null" json:"username"`
	PasswordHash string    `gorm:"column:passwordhash;size:255;not null" json:"-"`
	Email        string    `gorm:"column:email;size:100;unique" json:"email"`
	RoleID       uint      `gorm:"column:roleid;not null" json:"roleId"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
}

// Role модель
type Role struct {
	RoleID      uint         `gorm:"column:roleid;primaryKey;autoIncrement" json:"roleId"`
	RoleName    string       `gorm:"column:rolename;size:50;unique;not null" json:"roleName"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions"`
}

// Permission модель
type Permission struct {
	PermissionID   uint   `gorm:"column:permissionid;primaryKey;autoIncrement" json:"permissionId"`
	PermissionName string `gorm:"column:permissionname;size:100;unique;not null" json:"permissionName"`
	Description    string `gorm:"column:description" json:"description"`
}

// RolePermission модель для связи многие-ко-многим
type RolePermission struct {
	RoleID       uint `gorm:"column:roleid;primaryKey" json:"roleId"`
	PermissionID uint `gorm:"column:permissionid;primaryKey" json:"permissionId"`
}
