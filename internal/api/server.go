package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniciusmtsantos/golang-ecommerce/configs"
	"github.com/viniciusmtsantos/golang-ecommerce/internal/api/rest"
	"github.com/viniciusmtsantos/golang-ecommerce/internal/api/rest/handlers"
)

func StartServer(config configs.ApplicationConfig) {
	app := fiber.New()

	rh := &rest.RestHandler{
		App: app,
	}

	setupRoutes(rh)

	app.Listen(config.ServerPort)
}

func setupRoutes(rh *rest.RestHandler) {
	handlers.SetupUserRoutes(rh)
}
