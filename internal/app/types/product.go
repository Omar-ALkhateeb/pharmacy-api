package types

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Barcode   string  `gorm:"unique_index;not_null" csv:"barcode"`
	Name      string  `gorm:"not_null;index:name" csv:"name"`
	ExpiresIn int     `gorm:"not_null;index:expires_in" csv:"expires_in"`
	Price     float32 `gorm:"not_null" csv:"price"`
	Category  string  `gorm:"not_null" csv:"category"`
	StockIns  []StockIn
	StockOuts []StockOut
}

// for view and export
type ProductInView struct {
	ID              uint
	Barcode         string
	Name            string
	CurrentQuantity int
	Category        string
	ExpiresIn       int
	Price           float32
}
