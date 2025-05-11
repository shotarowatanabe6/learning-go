package main

import (
	"flag"
	"log"

	"restful-api-server/internal/api"
)

func main() {
	configFilePath := flag.String("config", "config/local.tml", "config file path")
	flag.Parse()

	config, err := api.NewConfig(*configFilePath)
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	api := api.NewAPI(config)
	api.Run()
}
