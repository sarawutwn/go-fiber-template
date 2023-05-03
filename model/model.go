package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID    string         `gorm:"type:uuid;primaryKey" json:"user_id"`
	Firstname string         `json:"firstname" validate:"required"`
	Lastname  string         `json:"lastname" validate:"required"`
	Username  string         `json:"username" validate:"required"`
	Password  string         `json:"password" validate:"required"`
	Status    bool           `json:"status" gorm:"default:true"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	RoleID    string         `json:"role_id"`
	Role      Role
}

type Role struct {
	RoleID          string            `gorm:"type:uuid;primaryKey" json:"role_id"`
	RoleName        string            `json:"role_name" validate:"required"`
	RoleDisplayName string            `json:"role_display_name" validate:"required"`
	RoleDescription string            `json:"role_description" validate:"required"`
	Status          bool              `json:"status" gorm:"default:true"`
	CreatedAt       *time.Time        `json:"created_at"`
	UpdatedAt       *time.Time        `json:"updated_at"`
	Users           []User            `gorm:"foreignKey:RoleID"`
	PermissionGroup []PermissionGroup `gorm:"foreignKey:RoleID"`
}

type Permission struct {
	PermissionID          string                `gorm:"type:uuid;primaryKey" json:"permission_id"`
	PermissionName        string                `json:"permission_name" validate:"required"`
	PermissionDisplayName string                `json:"permission_display_name" validate:"required"`
	PermissionDescription string                `json:"permission_description" validate:"required"`
	Status                bool                  `json:"status" gorm:"default:true"`
	CreatedAt             *time.Time            `json:"created_at"`
	UpdatedAt             *time.Time            `json:"updated_at"`
	PermissionComponent   []PermissionComponent `gorm:"foreignKey:PermissionID"`
	PermissionGroup       []PermissionGroup     `gorm:"foreignKey:PermissionID"`
}

type PermissionComponent struct {
	PermissionComponentID          string            `gorm:"type:uuid;primaryKey" json:"permission_component_id"`
	PermissionComponentName        string            `json:"permission_component_name" validate:"required"`
	PermissionComponentDisplayName string            `json:"permission_component_display_name" validate:"required"`
	PermissionComponentDescription string            `json:"permission_component_description" validate:"required"`
	Status                         bool              `json:"status" gorm:"default:true"`
	CreatedAt                      *time.Time        `json:"created_at"`
	UpdatedAt                      *time.Time        `json:"updated_at"`
	PermissionID                   string            `json:"permission_id"`
	Permission                     Permission        `json:"permission"`
	PermissionGroup                []PermissionGroup `gorm:"foreignKey:PermissionComponentID"`
}

type PermissionGroup struct {
	PermissionGroupID     string              `gorm:"type:uuid;primaryKey" json:"permission_group_id"`
	Activate              bool                `json:"avtivates"`
	CreatedAt             *time.Time          `json:"created_at"`
	UpdatedAt             *time.Time          `json:"updated_at"`
	RoleID                string              `json:"role_id"`
	PermissionID          string              `json:"permission_id"`
	PermissionComponentID string              `json:"permission_component_id"`
	Role                  Role                `json:"role"`
	Permission            Permission          `json:"permission"`
	PermissionComponent   PermissionComponent `json:"permission_component"`
}
