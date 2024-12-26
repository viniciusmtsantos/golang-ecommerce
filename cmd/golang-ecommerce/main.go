package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/viniciusmtsantos/golang-ecommerce/configs"
)

func main() {
	fmt.Println("Hello")

	app := fiber.New()

	configs.LoadSettings()

	app.Listen("localhost:9000")
}
