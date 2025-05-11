package api

import (
	"net/http"
	"restful-api-server/internal/api/handler"
	"restful-api-server/pkg/logger"
)

type API struct {
	config *Config
	logger *logger.Logger
	server *http.Server
}

func NewAPI(config *Config) *API {
	return &API{
		config: config,
		logger: logger.NewLogger().WithLevel(config.Logger.Level),
		server: handler.NewServer(config.Server.Port),
	}
}

func (a *API) Run() {
	a.logger.Infof("starting server on port %d", a.config.Server.Port)
	err := a.server.ListenAndServe()
	if err != nil {
		a.logger.Errorf("server error: %v", err)
	}
}
