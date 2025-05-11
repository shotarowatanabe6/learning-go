package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) GetProducts(c *gin.Context) {
	products, err := r.service.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}
