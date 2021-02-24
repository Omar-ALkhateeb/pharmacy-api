package handlers

import (
	"github.com/Omar-ALkhateeb/pharm-inventory/configs/database"
	"github.com/Omar-ALkhateeb/pharm-inventory/internal/app/types"
	"github.com/gin-gonic/gin"
)

func ProductListOne(c *gin.Context) {
	db := database.DBConn
	var product types.Product

	if err := db.Where("barcode = ?", c.Param("barcode")).Preload("StockIns").
		Preload("StockOuts").First(&product).Error; err != nil {
		// fmt.Println(err)
		c.String(404, "Product does not exist")
		return
	}

	currentQuantity := calculateCurrentQuantity(product)
	productShow := types.ProductInView{
		Barcode:         product.Barcode,
		Name:            product.Name,
		CurrentQuantity: currentQuantity,
	}

	c.JSON(200, productShow)
}
