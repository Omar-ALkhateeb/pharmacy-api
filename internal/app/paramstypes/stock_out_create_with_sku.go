package paramstypes

type StockOutWithParams struct {
	Barcode         string  `json:"barcode"`
	PricePerProduct float32 `json:"price_per_product"`
	Note            string  `json:"note"`
	Quantity        int     `json:"quantity"`
	Currency        string  `json:"currency"`
}
