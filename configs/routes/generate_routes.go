package routes

import (
	"github.com/gin-gonic/gin"
)

func GenerateRoutes() *gin.Engine {
	web := gin.Default()
	web.GET("/ping", HandlePing)

	return web
}

func HandlePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
