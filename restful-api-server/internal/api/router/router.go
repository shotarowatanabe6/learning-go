package router

import (
	"net/http"

	"restful-api-server/pkg/logger"

	"github.com/gin-gonic/gin"
)

func NewRouter(logger *logger.Logger) *gin.Engine {
	router := gin.Default()
	router.GET("/", helloWorld)

	return router
}

func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}
