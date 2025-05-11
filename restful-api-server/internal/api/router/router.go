package router

import (
	"net/http"
	"restful-api-server/internal/service"
	"restful-api-server/pkg/logger"

	"github.com/gin-gonic/gin"
)

type Router struct {
	service *service.Service
	engine  *gin.Engine
}

func NewRouter(logger *logger.Logger, svc *service.Service) *gin.Engine {
	router := gin.Default()

	r := &Router{
		service: svc,
		engine:  router,
	}

	router.GET("/", helloWorld)

	// product
	router.GET("/products", r.GetProducts)

	return router
}

func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}
