package handlers

import (
	"github.com/dwahyudi/inventory/configs/database"
	"github.com/dwahyudi/inventory/internal/app/paramstypes"
	"github.com/dwahyudi/inventory/internal/app/types"
	"github.com/gin-gonic/gin"
	"time"
)

func CreateStockOut(c *gin.Context) {
	db := database.DBConn
	var responseCode int
	responseMessage := ""

	var stockOutParams paramstypes.StockOutWithParams

	if c.ShouldBind(&stockOutParams) == nil {
		var product types.Product
		if err := db.Where("sku = ?", stockOutParams.Sku).First(&product).Error; err != nil {
			c.String(422, "Product does not exist")
			return
		} else {
			newQuantity := product.CurrentQuantity - stockOutParams.Quantity
			db.Model(&product).Update(types.Product{CurrentQuantity: newQuantity})
		}

		if errors := db.Create(&types.StockOut{
			PricePerProduct: stockOutParams.PricePerProduct,
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
