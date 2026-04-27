package entities

import "github.com/muhamadairul/string-parser-api/app/models"

// ParsedResult is the parsed string DB record.
type ParsedResult struct {
	ID   uint   `gorm:"primarykey" json:"id"`
	Name string `gorm:"type:char(30);not null" json:"name"`
	Age  string `gorm:"type:char(3);not null" json:"age"`
	City string `gorm:"type:char(20);not null" json:"city"`
	models.Model
}

// TableName overrides the default table name.
func (ParsedResult) TableName() string {
	return "parsed_results"
}
