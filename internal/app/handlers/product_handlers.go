package handlers

import (
	"github.com/Omar-ALkhateeb/pharm-inventory/configs/database"
	"github.com/Omar-ALkhateeb/pharm-inventory/internal/app/paramstypes"
	"github.com/Omar-ALkhateeb/pharm-inventory/internal/app/types"
	"github.com/Omar-ALkhateeb/pharm-inventory/internal/services"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product paramstypes.ProductParams

	var responseCode int
	responseMessage := ""
	if c.ShouldBind(&product) == nil {

		var convertionRate float32 = 1.0
		if len(product.Currency) > 0 {
			convertionRate = services.ConvertCurrency(product.Currency)
		}

		var category string
		if len(product.Category) > 0 {
			category = product.Category
		} else {
			category = "All"
		}

		if errors := database.DBConn.Create(&types.Product{
			Name:      product.Name,
			Barcode:   product.Barcode,
			ExpiresIn: product.ExpiresIn,
			Price:     product.Price * convertionRate,
			Category:  category,
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
