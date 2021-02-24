package types

import (
	"time"

	"github.com/jinzhu/gorm"
)

type StockOut struct {
	gorm.Model
	Time            time.Time `gorm:"not_null"`
	ProductId       int       `gorm:"not_null"`
	Product         Product
	PricePerProduct float32 `gorm:"not_null"`
	Note            string
	Quantity        int
}

type StockOutView struct {
	Time            time.Time
	Product         Product
	PricePerProduct float32
	Note            string
	Quantity        int
}
