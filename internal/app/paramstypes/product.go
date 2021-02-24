package paramstypes

type UpdateProductParam struct {
	Name string `json:"name"`
}

type ProductParams struct {
	Barcode   string  `json:"barcode"`
	Name      string  `json:"name"`
	ExpiresIn int     `json:"expires_in"`
	Price     float32 `json:"price"`
	Currency  string  `json:"currency"`
	Category  string  `json:"category"`
}
