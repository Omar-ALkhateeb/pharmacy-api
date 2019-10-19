package database

import "github.com/dwahyudi/inventory/internal/app/types"

func Migrate() {
	// Using auto migration feature from GORM framework.
	DBConn.AutoMigrate(
		&types.Product{},
		&types.StockIn{},
		&types.StockOut{})
}
