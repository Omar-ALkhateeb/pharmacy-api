package handlers

import (
	"time"

	"github.com/Omar-ALkhateeb/pharm-inventory/configs/database"
	"github.com/Omar-ALkhateeb/pharm-inventory/internal/app/paramstypes"
	"github.com/Omar-ALkhateeb/pharm-inventory/internal/app/types"
	"github.com/Omar-ALkhateeb/pharm-inventory/internal/services"
	"github.com/gin-gonic/gin"
)

func CreateStockOut(c *gin.Context) {
	db := database.DBConn
	var responseCode int
	responseMessage := ""

	var stockOutParams paramstypes.StockOutWithParams

	if c.ShouldBind(&stockOutParams) == nil {
		var product types.Product
		if err := db.Where("barcode = ?", stockOutParams.Barcode).First(&product).Error; err != nil {
			c.String(422, "Product does not exist")
			return
		}

		var convertionRate float32 = 1.0
		if len(stockOutParams.Currency) > 0 {
			convertionRate = services.ConvertCurrency(stockOutParams.Currency)
		}

		if errors := db.Create(&types.StockOut{
			PricePerProduct: stockOutParams.PricePerProduct * convertionRate,
			Quantity:        stockOutParams.Quantity,
			Product:         product,
			Note:            stockOutParams.Note,
			Time:            time.Now(),
		}).GetErrors(); len(errors) > 0 {
			responseCode = 422
			for _, err := range errors {
				responseMessage = responseMessage + ", " + err.Error()
			}
		} else {
			responseCode = 201
			responseMessage = "Created"
		}
	}

	c.String(responseCode, responseMessage)
}
