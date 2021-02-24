package handlers

import (
	"github.com/Omar-ALkhateeb/pharm-inventory/configs/database"
	"github.com/Omar-ALkhateeb/pharm-inventory/internal/app/types"
	"github.com/gin-gonic/gin"
)

func StockOutsList(c *gin.Context) {
	db := database.DBConn
	var stockOuts []types.StockOut

	db.Preload("Product").Find(&stockOuts)

	var StockOutsList []types.StockOutView

	for _, stockOut := range stockOuts {
		stockOutShow := types.StockOutView{
			Time:            stockOut.Time,
			PricePerProduct: stockOut.PricePerProduct,
			Product:         stockOut.Product,
			Note:            stockOut.Note,
			Quantity:        stockOut.Quantity,
		}

		StockOutsList = append(StockOutsList, stockOutShow)
	}

	c.JSON(200, StockOutsList)
}
