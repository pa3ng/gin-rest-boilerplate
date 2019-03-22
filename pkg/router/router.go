package app

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("/", Index)
		v1.GET("/health", Health)
	}

	return router
}
