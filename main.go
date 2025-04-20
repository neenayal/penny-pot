package main

import (
	"github.com/gin-gonic/gin"
	"penny-pot/routes"
)

func main() {
	r := gin.Default()
	
	routes.RegisterIndexRoutes(r)
	routes.RegisterCacheRoutes(r)

	r.Run(":8080") // start server
}
