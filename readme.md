# Inventory Project

Record inventory/stocks for products.

* Record product data with current quantity.
* Record stock ins with ordered and received quantity.
* Record stock out.
* Generate csv report for inventory values.
* Generate csv report for stock out.

Requirements:

* Golang 1.12

Dependencies:

* go get -u github.com/gin-gonic/gin
* go get -v github.com/mattn/go-sqlite3
* go get -u github.com/jinzhu/gorm

How to run:

* Install Golang 1.12
* Set $GOPATH variable
* go get -u github.com/dwahyudi/inventory
* go build github.com/dwahyudi/inventory
* run ./inventory

Project structure:

* `internal/` handles app code.
  1. `app/types` contains models/domains related to database.
  2. `app/handlers` contains handlers for handling business logic.
  3. `services` contains generic reusable services.
* `configs/` contains configurations such as simple database connection and migration and routings.