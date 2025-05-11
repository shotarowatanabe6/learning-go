package api

import (
	"fmt"
	"net/http"
	"restful-api-server/internal/api/handler"
)

type API struct {
	config *Config
	server *http.Server
}

func NewAPI(config *Config) *API {
	return &API{
		config: config,
		server: handler.NewServer(config.Server.Port),
	}
}

func (a *API) Run() error {
	if err := a.server.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}
