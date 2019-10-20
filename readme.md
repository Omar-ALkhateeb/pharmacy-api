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