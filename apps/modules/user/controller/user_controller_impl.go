package controller

import (
	"go-fiber-starter/apps/domain"
	"go-fiber-starter/apps/helpers"
	"go-fiber-starter/apps/modules/user/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserControllerImpl struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &UserControllerImpl{service: service}
}

func (c *UserControllerImpl) FindAll(ctx *fiber.Ctx) error {
	users, err := c.service.FindAll(ctx.Context())
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to fetch users")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, users, "Users fetched successfully")
}

func (c *UserControllerImpl) FindByID(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid user ID")
	}
	user, err := c.service.FindByID(ctx.Context(), id)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusNotFound, "User not found")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, user, "User fetched successfully")
}

func (c *UserControllerImpl) Create(ctx *fiber.Ctx) error {
	var user domain.User
	if err := ctx.BodyParser(&user); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request payload")
	}
	createdUser, err := c.service.Create(ctx.Context(), user)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to create user")
	}
	return helpers.WriteJSON(ctx, fiber.StatusCreated, createdUser, "User created successfully")
}

func (c *UserControllerImpl) Update(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid user ID")
	}
	var user domain.User
	if err := ctx.BodyParser(&user); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request payload")
	}
	user.ID = id
	updatedUser, err := c.service.Update(ctx.Context(), user)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to update user")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, updatedUser, "User updated successfully")
}

func (c *UserControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid user ID")
	}
	if err := c.service.Delete(ctx.Context(), id); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to delete user")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, nil, "User deleted successfully")
}
