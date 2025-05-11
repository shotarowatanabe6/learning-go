package api

import (
	"fmt"

	"restful-api-server/internal/api/handler"

	"github.com/gin-gonic/gin"
)

type API struct {
	config  *Config
	handler *gin.Engine
}

func NewAPI(config *Config) *API {
	return &API{
		config:  config,
		handler: handler.NewHandler(),
	}
}

func (a *API) Run() {
	a.handler.Run(fmt.Sprintf(":%d", a.config.Server.Port))
}
