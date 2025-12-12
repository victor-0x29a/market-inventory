package database

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID                uint      `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Title             string    `gorm:"size:64;not null" json:"title"`
	Description       *string   `gorm:"size:256" json:"description"`
	Price             int64     `gorm:"not null" json:"price"`
	InventoryQuantity int64     `gorm:"not null;default:0" json:"inventory_quantity"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
