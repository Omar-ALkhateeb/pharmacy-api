package handlers

import (
	"fmt"

	"github.com/Omar-ALkhateeb/pharm-inventory/configs/database"
	"github.com/Omar-ALkhateeb/pharm-inventory/internal/app/reporttypes"
	"github.com/Omar-ALkhateeb/pharm-inventory/internal/app/types"
	"github.com/Omar-ALkhateeb/pharm-inventory/internal/services"
	"github.com/gin-gonic/gin"
)

func InventoryValuationSummary(c *gin.Context) {
	db := database.DBConn
	var inventoryValuationSummary reporttypes.InventoryValuationSummary

	var productCount int
	db.Model(&types.Product{}).Count(&productCount)

	var totalStockIn reporttypes.SumResult
	db.Model(&types.StockIn{}).Select("sum(received_quantity) as sum").Scan(&totalStockIn)

	fmt.Println(totalStockIn)

	var totalStockOut reporttypes.SumResult
	db.Model(&types.StockOut{}).Select("sum(quantity) as sum").Scan(&totalStockOut)

	fmt.Println(totalStockOut)
	productTotalQuantity := totalStockIn.Sum - totalStockOut.Sum

	inventoryValuations := services.InventoryValuationCalculate()
	var totalValuation float32 = 0.0
	for _, iv := range inventoryValuations {
		totalValuation += iv.ProductTotalPurchasePrice
	}

	inventoryValuationSummary.ProductSkuCount = productCount
	inventoryValuationSummary.ProductTotalQuantity = productTotalQuantity
	inventoryValuationSummary.TotalValuation = totalValuation

	//json, err := json.Marshal(inventoryValuationSummary)

	c.JSON(200, inventoryValuationSummary)
}
