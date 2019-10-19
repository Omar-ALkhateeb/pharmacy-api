package handlers

import (
	"github.com/dwahyudi/inventory/configs/database"
	"github.com/dwahyudi/inventory/internal/app/types"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product types.Product
	if c.ShouldBind(&product) == nil {
		database.DBConn.Create(&types.Product{
			Name:            product.Name,
			Sku:             product.Sku,
			CurrentQuantity: product.CurrentQuantity})
	}

	c.String(201, "Created")
}
