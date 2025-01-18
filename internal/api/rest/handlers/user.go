package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/viniciusmtsantos/golang-ecommerce/internal/api/rest"
	"github.com/viniciusmtsantos/golang-ecommerce/internal/dto"
	"github.com/viniciusmtsantos/golang-ecommerce/internal/repository"
	"github.com/viniciusmtsantos/golang-ecommerce/internal/service"
)

type UserHandler struct {
	svc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {

	app := rh.App

	svc := service.UserService{
		Repo: repository.NewUserRepository(rh.DB),
		Auth: rh.Auth,
	}

	handler := UserHandler{
		svc: svc,
	}

	pubRoutes := app.Group("/users")
	pubRoutes.Post("/register", handler.Register)
	pubRoutes.Post("/login", handler.Login)

	pvtRoutes := pubRoutes.Group("/", rh.Auth.Authorize)
	pvtRoutes.Get("/verify", handler.GetVerificationCode)
	pvtRoutes.Post("/verify", handler.Verify)
	pvtRoutes.Post("/profile", handler.CreateProfile)
	pvtRoutes.Get("/profile", handler.GetProfile)

	pvtRoutes.Post("/cart", handler.AddToCart)
	pvtRoutes.Get("/cart", handler.GetCart)
	pvtRoutes.Get("/order", handler.GetOrders)
	pvtRoutes.Get("/order/:id", handler.GetOrder)

}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {

	user := dto.UserSignUp{}

	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide valid inputs",
		})
	}

	token, err := h.svc.SignUp(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "error on signup",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "register",
		"token":   token,
	})
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	loginInput := dto.UserLogin{}

	err := ctx.BodyParser(&loginInput)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide valid inputs",
		})
	}

	token, err := h.svc.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "please provide correct user id password",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login",
		"token":   token,
	})
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {

	user := h.svc.Auth.GetCurrentUser(ctx)

	code, err := h.svc.GetVerificationCode(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "unable to generate verification code",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get verification",
		"data":    code,
	})
}

func (h *UserHandler) Verify(ctx *fiber.Ctx) error {

	var req dto.VerificationCodeInput

	user := h.svc.Auth.GetCurrentUser(ctx)

	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide a valid input",
		})
	}

	err = h.svc.VerifyCode(user.ID, req.Code)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "verified successfully",
	})
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create profile",
	})
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {

	user := h.svc.Auth.GetCurrentUser(ctx)

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get profile",
		"user":    user,
	})
}

func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "add to cart",
	})
}

func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get cart",
	})
}

func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get orders",
	})
}

func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get order",
	})
}
