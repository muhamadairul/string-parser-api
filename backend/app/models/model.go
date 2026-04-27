package models

import (
	"time"

	"gorm.io/gorm"
)

// Model is the base model embedded in all entities.
type Model struct {
	CreatedAt time.Time      `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time      `gorm:"default:current_timestamp" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
