package app

import (
	"github.com/gin-gonic/gin"
)

func Root(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Gin Boilerplate API v1"})

}

func Health(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}
