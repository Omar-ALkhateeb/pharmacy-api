package reporttypes

import "time"

type SalesReport struct {
	SalesNote               string
	Time                    time.Time
	ProductSku              string
	ProductName             string
	Quantity                int
	SellPricePerProduct     float32
	TotalPricePerProduct    float32
	BuyPricePerProduct      float32
	TotalBuyPricePerProduct float32
	ProfitOrLoss            float32
}
