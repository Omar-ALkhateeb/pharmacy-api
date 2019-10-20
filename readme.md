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
  4. `app/paramstypes` contains structs to handle incoming web parameters.
  5. `app/reporttypes` contains structs to handle report data modelling.
* `configs/` contains configurations such as simple database connection and migration and routings.

Request Samples:
(Check `configs/routes` for more detail.)

Create product (user must create product before creating stock in and stock out).

```
POST /v1/products HTTP/1.1
Host: localhost:8080
Content-Type: application/json
cache-control: no-cache

{
	"sku": "A0",
	"name": "drawing book"
}
```

Create Stock-in, must provide SKU of already made product.

```
POST /v1/stock_ins HTTP/1.1
Host: localhost:8080
Content-Type: application/json
cache-control: no-cache

{
	"sku": "A0",
	"price_per_product": 11000.0,
	"transaction_number": "T0",
	"ordered_quantity": 100,
	"received_quantity": 100
}
```

Create Stock-out, also must provide valid product SKU.

```
POST /v1/stock_outs HTTP/1.1
Host: localhost:8080
Content-Type: application/json
cache-control: no-cache

{
	"sku": "A0",
	"price_per_product": 15000.0,
	"quantity": 80
}
```

Get Inventory Valuation Report in CSV

```
GET /v1/reports/inventory_valuation.csv HTTP/1.1
Host: localhost:8080
cache-control: no-cache
```