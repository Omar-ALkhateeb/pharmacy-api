package database

import models "github.com/dwahyudi/inventory/internal/app"

func Migrate() {
	// Using auto migration feature from GORM framework.
	DBConn.AutoMigrate(
		&models.Product{},
		&models.StockIn{},
		&models.StockOut{})
}
