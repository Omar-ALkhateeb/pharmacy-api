package types

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Sku             string `gorm:"unique_index;not_null"`
	Name            string `gorm:"not_null;index:name"`
	CurrentQuantity int    `gorm:"not_null;default:0"`
}
