package handlers

import (
	"strconv"

	"github.com/Omar-ALkhateeb/pharm-inventory/configs/database"
	"github.com/Omar-ALkhateeb/pharm-inventory/internal/app/types"
	"github.com/gin-gonic/gin"
)

func ProductList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "3"))
	db := database.DBConn
	var products []types.Product

	db.Offset(page * limit).Limit(limit).Preload("StockIns").Preload("StockOuts").
		Find(&products)

	var productsList []types.ProductInView

	for _, product := range products {
		currentQuantity := calculateCurrentQuantity(product)
		productShow := types.ProductInView{
			Barcode:         product.Barcode,
			Name:            product.Name,
			Category:        product.Category,
			CurrentQuantity: currentQuantity,
			ExpiresIn:       product.ExpiresIn,
		}

		productsList = append(productsList, productShow)
	}

	c.JSON(200, productsList)
}

func calculateCurrentQuantity(product types.Product) int {
	totalStockInsQuantity := 0
	totalStockOutsQuantity := 0

	for _, stockIn := range product.StockIns {
		totalStockInsQuantity += stockIn.ReceivedQuantity
	}

	for _, stockOut := range product.StockOuts {
		totalStockOutsQuantity += stockOut.Quantity
	}

	return totalStockInsQuantity - totalStockOutsQuantity
}