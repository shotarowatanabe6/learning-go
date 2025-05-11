package api

import (
	"fmt"
	"restful-api-server/internal/api/router"
	"restful-api-server/internal/config"
	"restful-api-server/internal/service"
	"restful-api-server/pkg/logger"

	"github.com/gin-gonic/gin"
)

type API struct {
	config  *config.Config
	logger  *logger.Logger
	router  *gin.Engine
	service *service.Service
}

func NewAPI(config *config.Config) *API {
	logger := logger.NewLogger().WithLevel(config.Logger.Level)
	svc := service.NewService(config)
	router := router.NewRouter(logger, svc)

	return &API{
		config:  config,
		logger:  logger,
		router:  router,
		service: svc,
	}
}

func (a *API) Run() {
	a.logger.Infof("starting server on port %d", a.config.Server.Port)
	err := a.router.Run(fmt.Sprintf(":%d", a.config.Server.Port))
	if err != nil {
		a.logger.Errorf("server error: %v", err)
	}
}
