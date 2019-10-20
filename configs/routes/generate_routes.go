package routes

import (
	"github.com/dwahyudi/inventory/internal/app/handlers"
	"github.com/gin-gonic/gin"
)

func GenerateRoutes() *gin.Engine {
	web := gin.Default()
	web.GET("/ping", handlers.HandlePing)

	web.POST("v1/products", handlers.CreateProduct)
	web.PATCH("v1/products/:id", handlers.UpdateProduct)
	web.DELETE("v1/products/:id", handlers.DeleteProduct)

	web.POST("v1/stock_ins", handlers.CreateStockIn)

	web.POST("v1/stock_outs", handlers.CreateStockOut)

	web.GET("v1/reports/inventory_valuation.csv", handlers.InventoryValuation)
	web.GET("v1/reports/sales_report.csv", handlers.SalesReport)

	return web
}
