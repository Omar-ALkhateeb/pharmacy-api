package main

import (
	"log"

	"github.com/Omar-ALkhateeb/pharm-inventory/configs/database"
	"github.com/Omar-ALkhateeb/pharm-inventory/configs/routes"
)

func main() {
	database.Prepare()

	web := routes.GenerateRoutes()

	log.Fatal(web.Run())
}
