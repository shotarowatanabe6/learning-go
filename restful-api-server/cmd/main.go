package main

import (
	"flag"
	"log"

	"restful-api-server/internal/api"
	"restful-api-server/internal/config"
)

func main() {
	// config
	configFilePath := flag.String("config", "config/local.tml", "config file path")
	flag.Parse()

	config, err := config.NewConfig(*configFilePath)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
		return
	}

	// run
	api := api.NewAPI(config)
	api.Run()
}
