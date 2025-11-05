package controller

import (
	"mou-be/apps/domain"
	"mou-be/apps/helpers"
	"mou-be/apps/modules/test/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TestControllerImpl struct {
	service service.TestService
}

func NewTestController(service service.TestService) TestController {
	return &TestControllerImpl{service: service}
}

func (c *TestControllerImpl) FindAll(ctx *fiber.Ctx) error {
	items, err := c.service.FindAll(ctx.Context())
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to fetch tests")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, items, "Tests fetched successfully")
}

func (c *TestControllerImpl) FindByID(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid test ID")
	}
	item, err := c.service.FindByID(ctx.Context(), id)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusNotFound, "Test not found")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, item, "Test fetched successfully")
}

func (c *TestControllerImpl) Create(ctx *fiber.Ctx) error {
	var item domain.Test
	if err := ctx.BodyParser(&item); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request payload")
	}
	created, err := c.service.Create(ctx.Context(), item)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to create test")
	}
	return helpers.WriteJSON(ctx, fiber.StatusCreated, created, "Test created successfully")
}

func (c *TestControllerImpl) Update(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid test ID")
	}
	var item domain.Test
	if err := ctx.BodyParser(&item); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request payload")
	}
	item.ID = id
	updated, err := c.service.Update(ctx.Context(), item)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to update test")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, updated, "Test updated successfully")
}

func (c *TestControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid test ID")
	}
	if err := c.service.Delete(ctx.Context(), id); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to delete test")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, nil, "Test deleted successfully")
}
