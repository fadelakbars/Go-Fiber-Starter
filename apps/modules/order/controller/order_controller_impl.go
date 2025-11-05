package controller

import (
	"mou-be/apps/domain"
	"mou-be/apps/helpers"
	"mou-be/apps/modules/order/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type OrderControllerImpl struct {
	service service.OrderService
}

func NewOrderController(service service.OrderService) OrderController {
	return &OrderControllerImpl{service: service}
}

func (c *OrderControllerImpl) FindAll(ctx *fiber.Ctx) error {
	limit := ctx.QueryInt("limit", 10)
	if limit <= 0 {
		limit = 10
	}
	offset := ctx.QueryInt("offset", 0)
	if offset < 0 {
		offset = 0
	}
	startDate := ctx.Query("startDate", "")
	endDate := ctx.Query("endDate", "")
	boothIDParam := ctx.Query("search", "")
	var boothID uuid.UUID
	var err error
	if boothIDParam != "" {
		boothID, err = uuid.Parse(boothIDParam)
		if err != nil {
			return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid booth ID")
		}
	}

	filters := map[string]interface{}{
		"limit":     limit,
		"offset":    offset,
		"startDate": startDate,
		"endDate":   endDate,
	}
	if boothIDParam != "" {
		filters["boothID"] = boothID
	}

	orders, total, totalPrice, err := c.service.FindAll(ctx.Context(), filters)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to fetch orders")
	}
	response := fiber.Map{
		"orders":      orders,
		"total":       total,
		"total_price": totalPrice,
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, response, "Orders fetched successfully")
}

func (c *OrderControllerImpl) FindByID(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid order ID")
	}
	order, err := c.service.FindByID(ctx.Context(), id)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusNotFound, "Order not found")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, order, "Order fetched successfully")
}

func (c *OrderControllerImpl) Create(ctx *fiber.Ctx) error {
	var order domain.Order
	if err := ctx.BodyParser(&order); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request payload")
	}
	createdOrder, err := c.service.Create(ctx.Context(), order)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to create order")
	}
	return helpers.WriteJSON(ctx, fiber.StatusCreated, createdOrder, "Order created successfully")
}

func (c *OrderControllerImpl) Update(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid order ID")
	}
	var order domain.Order
	if err := ctx.BodyParser(&order); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request payload")
	}
	order.ID = id
	updatedOrder, err := c.service.Update(ctx.Context(), order)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to update order")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, updatedOrder, "Order updated successfully")
}

func (c *OrderControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid order ID")
	}
	if err := c.service.Delete(ctx.Context(), id); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to delete order")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, nil, "Order deleted successfully")
}
