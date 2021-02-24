package types

import (
	"time"

	"github.com/jinzhu/gorm"
)

type StockIn struct {
	gorm.Model
	StartTime         time.Time `gorm:"not_null"` // manufacture time
	EndTime           time.Time `gorm:"not_null"` // expire time
	ProductId         int       `gorm:"not_null"`
	Product           Product
	PricePerProduct   float32 `gorm:"not_null"`
	TransactionNumber string  `gorm:"unique_index;not_null"`
	Note              string
	OrderedQuantity   int
	ReceivedQuantity  int
}

type StockInView struct {
	StartTime         time.Time // manufacture time
	EndTime           time.Time // expire time
	Product           Product
	PricePerProduct   float32
	TransactionNumber string
	Note              string
	OrderedQuantity   int
	ReceivedQuantity  int
	Expired           bool
}
