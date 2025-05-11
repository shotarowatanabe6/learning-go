package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandler() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	return router
}
