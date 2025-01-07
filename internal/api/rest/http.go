package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniciusmtsantos/golang-ecommerce/internal/helper"
	"gorm.io/gorm"
)

type RestHandler struct {
	App  *fiber.App
	DB   *gorm.DB
	Auth helper.Auth
}
