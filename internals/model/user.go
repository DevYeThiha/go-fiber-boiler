package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model        // Adds some metadata fields to the table
	ID         uint   `gorm:"primaryKey"` // Explicitly specify the type to be uuid
	UUID       string `json:"UUID"`       // Explicitly specify the type to be uuid
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}
