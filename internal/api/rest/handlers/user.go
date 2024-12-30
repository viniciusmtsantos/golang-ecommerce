package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/viniciusmtsantos/golang-ecommerce/internal/api/rest"
)

type User struct {
	// svc UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {

	app := rh.App

	handler := User{}

	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	app.Get("/verify", handler.GetVerificationCode)
	app.Post("/verify", handler.Verify)
	app.Post("/profile", handler.CreateProfile)
	app.Get("/profile", handler.GetProfile)

	app.Post("/cart", handler.AddToCart)
	app.Get("/cart", handler.GetCart)
	app.Get("/order", handler.GetOrders)
	app.Get("/order/:id", handler.GetOrder)

}

func (u *User) Register(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "register",
	})
}

func (u *User) Login(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login",
	})
}

func (u *User) GetVerificationCode(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get verification",
	})
}

func (u *User) Verify(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "verify",
	})
}

func (u *User) CreateProfile(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create profile",
	})
}

func (u *User) GetProfile(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get profile",
	})
}

func (u *User) AddToCart(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "add to cart",
	})
}

func (u *User) GetCart(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get cart",
	})
}

func (u *User) GetOrders(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get orders",
	})
}

func (u *User) GetOrder(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get order",
	})
}
