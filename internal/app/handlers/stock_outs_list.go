package handlers

import (
	"strconv"

	"github.com/Omar-ALkhateeb/pharm-inventory/configs/database"
	"github.com/Omar-ALkhateeb/pharm-inventory/internal/app/types"
	"github.com/gin-gonic/gin"
)

func StockOutsList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	db := database.DBConn
	var stockOuts []types.StockOut

	db.Offset(page * limit).Preload("Product").Find(&stockOuts)

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
