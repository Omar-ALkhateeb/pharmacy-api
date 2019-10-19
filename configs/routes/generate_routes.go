package routes

import (
	"github.com/dwahyudi/inventory/internal/app/handlers"
	"github.com/gin-gonic/gin"
)

func GenerateRoutes() *gin.Engine {
	web := gin.Default()
	web.GET("/ping", handlers.HandlePing)

	web.POST("v1/products", handlers.CreateProduct)

	return web
}
