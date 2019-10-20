package handlers

import (
	"github.com/dwahyudi/inventory/configs/database"
	"github.com/dwahyudi/inventory/internal/app/types"
	"github.com/gin-gonic/gin"
)

func DeleteProduct(c *gin.Context) {
	db := database.DBConn
	product := types.Product{}
	id := c.Param("id")

	// =============================================================================
	// VALIDATIONS
	// =============================================================================
	if err := db.First(&product, id).Error; err != nil {
		c.String(404, "Product Not Found")
		return
	}

	// =============================================================================

	db.Unscoped().Delete(&product)

	c.String(200, "Product Deleted")
}
