# Inventory Project

Record inventory/stocks for products.

- Record product data with current quantity.
- Record stock ins with ordered and received quantity.
- Record stock out.
- Generate csv report for inventory values.
- Generate csv report for sales report.
- mass import stock from csv files.
- use multiple currencies simultaneously without any errors.
- filter and search through products.
- pagination support.

## Requirements:

- Golang 1.12

## Dependencies:

- go get

## Project structure:

- `internal/` handles app code.
  1. `app/types` contains models/domains related to database.
  2. `app/handlers` contains handlers for handling business logic.
  3. `services` contains generic reusable services.
  4. `app/paramstypes` contains structs to handle incoming web parameters.
  5. `app/reporttypes` contains structs to handle report data modelling.
- `configs/` contains configurations such as simple database connection and migration and web-routings.

## Notes:

- Product must be created first. Barcode is unique.
- Transaction Number in stock-in is unique.
- User will reference product in stock-in and stock-out with product's Barcode. If product with such Barcode doesn't exist, response will be 422.
- Product can be updated, but only its name.
- Product can be destroyed (only if there is no stock-in and stock-out for that product).
- User may create "negative inventory" (can create stock-out without stock-in).
- Stock in and Stock out can be updated and deleted.
- Product current quantity is not stored in database (generated on request).
- all requests with prices can have an extra field (Currency) which our uom service will exchange into USD since this is the default currency for the rest of the system
- new currencies can be easily added through internal/services/uom.go

## Request Samples:

(Check `configs/routes` for more detail.)

Create product (user must create product before creating stock in and stock out).

```
POST /v1/products HTTP/1.1
Host: localhost:8080
Content-Type: application/json
cache-control: no-cache

{
	"Barcode": "A0",
	"name": "drawing book",
	"Category": "All",
	"ExpiresIn": 200,
	"Price": 40,
	"currency": "IQD"
}
```

Get Products List

```
GET /v1/products?page=0&limit=5&name=draw HTTP/1.1
Host: localhost:8080
Content-Type: application/json
cache-control: no-cache
```

Create Stock-in, must provide Barcode of an already made product.

```
POST /v1/stock_ins HTTP/1.1
Host: localhost:8080
Content-Type: application/json
cache-control: no-cache

{
	"Barcode": "A0",
	"price_per_product": 11000.0,
	"transaction_number": "T0",
	"ordered_quantity": 100,
	"received_quantity": 100,
	"currency": "USD"
}
```

Create Stock-out, also must provide valid product Barcode.

```
POST /v1/stock_outs HTTP/1.1
Host: localhost:8080
Content-Type: application/json
cache-control: no-cache

{
	"Barcode": "A0",
	"price_per_product": 15000.0,
	"quantity": 80,
	"currency": "IQD"
}
```

Get Inventory Valuation Report in CSV

```
GET /v1/reports/inventory_valuation.csv HTTP/1.1
Host: localhost:8080
cache-control: no-cache
```

Get Sales Report (start date and end date is optional, supply none, either or both)

```
GET /v1/reports/sales_report.csv?start_date=2019-01-01&amp; end_date=2019-12-30 HTTP/1.1
Host: localhost:8080
cache-control: no-cache

```
