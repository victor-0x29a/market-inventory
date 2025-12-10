package database

import "time"

type Product struct {
	ID                uint    `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Title             string  `gorm:"size:64;not null" json:"title"`
	Description       *string `gorm:"size:256" json:"description"`
	Price             int64   `gorm:"not null" json:"price"`
	InventoryQuantity int64   `gorm:"not null;default:0" json:"inventory_quantity"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
