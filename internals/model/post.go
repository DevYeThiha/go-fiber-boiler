package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model        // Adds some metadata fields to the table
	ID         uint   `gorm:"primaryKey"` // Explicitly specify the type to be uuid
	UUID       string `json:"UUID"`       // Explicitly specify the type to be uuid
	Title      string `json:"Title"`
	Detail     string `json:"Detail"`
}
