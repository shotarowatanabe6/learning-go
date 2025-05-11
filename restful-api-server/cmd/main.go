package main

import (
	"flag"
	"log"

	"restful-api-server/internal/api"
)

func main() {
	// config
	configFilePath := flag.String("config", "config/local.tml", "config file path")
	flag.Parse()

	config, err := api.NewConfig(*configFilePath)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
		return
	}

	// run
	api := api.NewAPI(config)
	api.Run()
}
