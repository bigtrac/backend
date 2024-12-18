package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	// gorm.Model          // Adds some metadata fields to the table
	// Userid       uuid `json:"userid" gorm:"unique;not null"`
	ID         uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Userid     string         `json:"userid"`
	DeviceId   string         `json:"device_id" gorm:"unique;not null"`
	DeviceOs   string         `json:"device_os" gorm:"not null"`
	DeviceType string         `json:"device_type" gorm:"not null"`
	Name       string         `json:"name"`
	Active     bool           `json:"active"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
