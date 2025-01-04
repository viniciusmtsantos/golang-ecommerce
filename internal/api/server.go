package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/viniciusmtsantos/golang-ecommerce/configs"
	"github.com/viniciusmtsantos/golang-ecommerce/internal/api/rest"
	"github.com/viniciusmtsantos/golang-ecommerce/internal/api/rest/handlers"
	"github.com/viniciusmtsantos/golang-ecommerce/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config configs.ApplicationConfig) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection error %v\n", err)
	}

	log.Println("database connected")

	db.AutoMigrate(&domain.User{})

	rh := &rest.RestHandler{
		App: app,
		DB:  db,
	}

	setupRoutes(rh)

	app.Listen(config.ServerPort)
}

func setupRoutes(rh *rest.RestHandler) {
	handlers.SetupUserRoutes(rh)
}
