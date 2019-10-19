package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Sku             string `gorm:"unique_index;not_null"`
	Name            string `gorm:"not_null;index:name"`
	CurrentQuantity int    `gorm:"not_null"`
}

type StockIn struct {
	gorm.Model
	Time              time.Time `gorm:"not_null"`
	ProductId         int       `gorm:"not_null"`
	Product           Product
	PricePerProduct   float32 `gorm:"not_null"`
	TransactionNumber string  `gorm:"unique_index;not_null"`
	Note              string
}

type StockOut struct {
	gorm.Model
	Time            time.Time `gorm:"not_null"`
	ProductId       int       `gorm:"not_null"`
	Product         Product
	PricePerProduct float32 `gorm:"not_null"`
	Note            string
}
