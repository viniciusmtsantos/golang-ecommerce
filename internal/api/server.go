package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniciusmtsantos/golang-ecommerce/configs"
)

func StartServer(config configs.ApplicationConfig) {
	app := fiber.New()

	app.Listen(config.ServerPort)
}
