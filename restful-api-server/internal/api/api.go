package api

import (
	"fmt"
	"restful-api-server/internal/api/router"
	"restful-api-server/pkg/logger"

	"github.com/gin-gonic/gin"
)

type API struct {
	config *Config
	logger *logger.Logger
	router *gin.Engine
}

func NewAPI(config *Config) *API {
	logger := logger.NewLogger().WithLevel(config.Logger.Level)
	router := router.NewRouter(logger)

	return &API{
		config: config,
		logger: logger,
		router: router,
	}
}

func (a *API) Run() {
	a.logger.Infof("starting server on port %d", a.config.Server.Port)
	err := a.router.Run(fmt.Sprintf(":%d", a.config.Server.Port))
	if err != nil {
		a.logger.Errorf("server error: %v", err)
	}
}
