package paramstypes

type StockInCreateWithSku struct {
	PricePerProduct   float32 `json:"price_per_product"`
	TransactionNumber string  `json:"transaction_number"`
	Note              string  `json:"note"`
	Barcode           string  `json:"barcode"`
	OrderedQuantity   int     `json:"ordered_quantity"`
	ReceivedQuantity  int     `json:"received_quantity"`
	StartTime         string  `json:"starttime"` // manufacture time
	EndTime           string  `json:"endtime"`   // expire time
	Currency          string  `json:"currency"`
}
