package routes

import (
	"github.com/Omar-ALkhateeb/pharm-inventory/internal/app/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GenerateRoutes() *gin.Engine {
	web := gin.Default()
	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"*"}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	// Register the middleware
	web.Use(cors.New(corsConfig))
	web.GET("/ping", handlers.HandlePing)

	web.GET("v1/products", handlers.ProductList)
	web.GET("v1/products/:barcode", handlers.ProductListOne)
	web.POST("v1/products", handlers.CreateProduct)
	web.POST("v1/import", handlers.ProductImport)
	web.PATCH("v1/products/:id", handlers.UpdateProduct)
	web.DELETE("v1/products/:id", handlers.DeleteProduct)

	web.GET("v1/stock_ins", handlers.StockInsList)
	web.POST("v1/stock_ins", handlers.CreateStockIn)
	web.PATCH("v1/stock_ins/:id", handlers.UpdateStockIn)
	web.DELETE("v1/stock_ins/:id", handlers.DeleteStockIn)

	web.GET("v1/stock_outs", handlers.StockOutsList)
	web.POST("v1/stock_outs", handlers.CreateStockOut)
	web.PATCH("v1/stock_outs/:id", handlers.UpdateStockOut)
	web.DELETE("v1/stock_outs/:id", handlers.DeleteStockOut)

	web.GET("v1/reports/inventory_valuation_summary", handlers.InventoryValuationSummary)
	web.GET("v1/reports/sales_report_summary", handlers.SalesReportSummary)

	web.GET("v1/reports/inventory_valuation.csv", handlers.InventoryValuation)
	web.GET("v1/reports/inventory_valuation", handlers.InventoryValuationJson)
	web.GET("v1/reports/sales_report.csv", handlers.SalesReport)
	web.GET("v1/reports/sales_report", handlers.SalesReportJson)

	return web
}
