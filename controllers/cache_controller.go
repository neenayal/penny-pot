package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCache(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Cache value"})
}

func SetCache(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Cache value created"})
}
