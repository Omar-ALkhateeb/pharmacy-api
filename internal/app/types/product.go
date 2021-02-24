package types

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Barcode   string  `gorm:"unique_index;not_null"`
	Name      string  `gorm:"not_null;index:name"`
	ExpiresIn int     `gorm:"not_null;index:expires_in"`
	Price     float32 `gorm:"not_null"`
	Category  string  `gorm:"not_null"`
	StockIns  []StockIn
	StockOuts []StockOut
}

type ProductInView struct {
	Barcode         string
	Name            string
	CurrentQuantity int
	Price           float32
}
