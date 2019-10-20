package handlers

import (
	"github.com/dwahyudi/inventory/configs/database"
	"github.com/dwahyudi/inventory/internal/app/paramstypes"
	"github.com/dwahyudi/inventory/internal/app/types"
	"github.com/gin-gonic/gin"
	"time"
)

// First find if product by such SKU exists.
// If not return response 422.
//
// If product already exists, add the quantity.
func CreateStockIn(c *gin.Context) {
	db := database.DBConn
	var responseCode int
	responseMessage := ""
	var StockInParams paramstypes.StockInCreateWithSku

	if c.ShouldBind(&StockInParams) == nil {
		var product types.Product
		if err := db.Where("sku = ?", StockInParams.Sku).First(&product).Error; err != nil {
			c.String(422, "Product does not exist")
			return
		} else {
			newQuantity := product.CurrentQuantity + StockInParams.ReceivedQuantity
			db.Model(&product).Update(types.Product{CurrentQuantity: newQuantity})
		}

		if errors := db.Create(&types.StockIn{
			PricePerProduct:   StockInParams.PricePerProduct,
			TransactionNumber: StockInParams.TransactionNumber,
			OrderedQuantity:   StockInParams.OrderedQuantity,
			ReceivedQuantity:  StockInParams.ReceivedQuantity,
			Product:           product,
			Note:              StockInParams.Note,
			Time:              time.Now(),
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
