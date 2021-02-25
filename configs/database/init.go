package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func Init() {
	var err error
	DBConn, err = gorm.Open(sqlite.Open("inventory.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
