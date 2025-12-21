package database

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title             string  `gorm:"size:64;not null" json:"Title"`
	Description       *string `gorm:"size:256" json:"Description"`
	Price             int64   `gorm:"not null" json:"Price"`
	InventoryQuantity int64   `gorm:"not null;default:0" json:"InventoryQuantity"`
}

type DamageLog struct {
	gorm.Model
	ProductID uint    `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int     `gorm:"not null"`
	Reason    int     `gorm:"not null"`
}
