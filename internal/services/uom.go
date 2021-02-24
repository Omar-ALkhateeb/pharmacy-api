package services

// mapped to usd
var currencyExchange = map[string]float32{
	"USD": 1.0,
	"IQD": 0.000685,
}

func ConvertCurrency(currency string) float32 {
	return currencyExchange[currency]
}
