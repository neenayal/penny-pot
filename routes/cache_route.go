package routes

import (
	"github.com/gin-gonic/gin"
	"penny-pot/controllers"
)

func RegisterCacheRoutes(r *gin.Engine) {
	cacheGroup := r.Group("/cache")
	{
		cacheGroup.GET("/get", controllers.GetCache)
		cacheGroup.GET("/set", controllers.SetCache)
	}
}
