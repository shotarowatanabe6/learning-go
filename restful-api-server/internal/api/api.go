package api

import (
	"log"
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

func (a *API) Run() {
	log.Printf("starting server on port %d", a.config.Server.Port)
	err := a.server.ListenAndServe()
	if err != nil {
		log.Fatalf("server error: %v", err)
	}
}
