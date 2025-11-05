package controller

import (
	"mou-be/apps/helpers"
	"mou-be/apps/modules/auth/service"

	"github.com/gofiber/fiber/v2"
)

type AuthControllerImpl struct {
	service service.AuthService
}

func NewAuthController(service service.AuthService) AuthController {
	return &AuthControllerImpl{service: service}
}

func (c *AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.BodyParser(&req); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request payload")
	}
	user, err := c.service.Login(ctx.Context(), req.Username, req.Password)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusUnauthorized, "Invalid username or password")
	}

	err = c.service.ComparePassword(user.Password, req.Password)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusUnauthorized, "Invalid credentials")
	}

	token, err := c.service.GenerateAuthToken(user)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to generate auth token")
	}

	return helpers.WriteJSON(ctx, fiber.StatusOK, fiber.Map{
		"token": token,
		"user":  user,
	}, "Login successful")
}
