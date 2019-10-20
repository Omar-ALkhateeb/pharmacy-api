package services

import (
	"github.com/dwahyudi/inventory/configs/database"
	"github.com/dwahyudi/inventory/internal/app/reporttypes"
	"github.com/dwahyudi/inventory/internal/app/types"
)

// Get inventory valuations in order to get buy price.
// Then for each stock out, obtain product buy price, then calculate the profit/loss.
func SalesReportCalculate() []reporttypes.SalesReport {
	db := database.DBConn
	inventoryValuations := InventoryValuationCalculate()
	var salesReports []reporttypes.SalesReport
	var stockOuts []types.StockOut

	skuAndAvgBuyPrices := make(map[string]float32)
	for _, iv := range inventoryValuations {
		skuAndAvgBuyPrices[iv.ProductSku] = iv.ProductAvgPurchasePrice
	}

	db.Preload("Product").Find(&stockOuts)

	for _, stockOut := range stockOuts {
		var salesReport = reporttypes.SalesReport{}
		salesReport.SalesNote = stockOut.Note
		salesReport.Time = stockOut.Time
		salesReport.ProductSku = stockOut.Product.Sku
		salesReport.ProductName = stockOut.Product.Name
		salesReport.Quantity = stockOut.Quantity

		salesReport.SellPricePerProduct = stockOut.PricePerProduct
		totalSellPrice := stockOut.PricePerProduct * float32(stockOut.Quantity)
		salesReport.TotalSellPricePerProduct = totalSellPrice

		buyPricePerProduct := skuAndAvgBuyPrices[stockOut.Product.Sku]
		salesReport.BuyPricePerProduct = buyPricePerProduct
		totalBuyPrice := buyPricePerProduct * float32(stockOut.Quantity)
		salesReport.TotalSellPricePerProduct = totalBuyPrice

		salesReport.ProfitOrLoss = totalSellPrice - totalBuyPrice
		salesReports = append(salesReports, salesReport)
	}

	return salesReports
}
