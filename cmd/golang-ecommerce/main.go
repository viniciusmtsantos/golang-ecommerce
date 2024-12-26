package main

import (
	"log"

	"github.com/viniciusmtsantos/golang-ecommerce/configs"
	"github.com/viniciusmtsantos/golang-ecommerce/internal/api"
)

func main() {
	cfg, err := configs.SetupEnvironment()
	if err != nil {
		log.Fatalf("config file was not loaded correctly: %v\n", err)
	}

	api.StartServer(cfg)
}
