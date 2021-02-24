package handlers

import (
	"github.com/Omar-ALkhateeb/pharm-inventory/configs/database"
	"github.com/Omar-ALkhateeb/pharm-inventory/internal/app/types"
	"github.com/gin-gonic/gin"
)

func StockInsList(c *gin.Context) {
	db := database.DBConn
	var stockIns []types.StockIn

	db.Preload("Product").Find(&stockIns)

	var StockInsList []types.StockInView

	for _, stockIn := range stockIns {
		expired := !stockIn.EndTime.After(stockIn.StartTime)
		stockInShow := types.StockInView{
			StartTime:         stockIn.StartTime,
			EndTime:           stockIn.EndTime,
			Expired:           expired,
			PricePerProduct:   stockIn.PricePerProduct,
			Product:           stockIn.Product,
			TransactionNumber: stockIn.TransactionNumber,
			Note:              stockIn.Note,
			ReceivedQuantity:  stockIn.ReceivedQuantity,
			OrderedQuantity:   stockIn.OrderedQuantity,
		}

		StockInsList = append(StockInsList, stockInShow)
	}

	c.JSON(200, StockInsList)
}
