package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterIndexRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the API 👋",
		})
	})
	
	r.POST("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the API 👋",
		})
	})
}
