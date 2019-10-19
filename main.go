package main

import (
	"github.com/dwahyudi/inventory/configs/routes"
	"log"
)

func main() {
	web := routes.GenerateRoutes()

	log.Fatal(web.Run())
}
