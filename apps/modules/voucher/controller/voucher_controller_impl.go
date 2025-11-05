package controller

import (
	"mou-be/apps/domain"
	"mou-be/apps/helpers"
	"mou-be/apps/modules/voucher/service"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type VoucherControllerImpl struct {
	service service.VoucherService
}

func NewVoucherController(service service.VoucherService) VoucherController {
	return &VoucherControllerImpl{service: service}
}

func (c *VoucherControllerImpl) FindAll(ctx *fiber.Ctx) error {
	limit := ctx.QueryInt("limit", 10)
	offset := ctx.QueryInt("offset", 0)
	search := ctx.Query("search", "")
	voucherTemplateID := ctx.Query("voucher_template_id", "")
	validFromStr := ctx.Query("validFrom", "")
	validUntilStr := ctx.Query("validUntil", "")

	filters := map[string]interface{}{
		"limit":  limit,
		"offset": offset,
	}

	if search != "" {
		filters["search"] = search
	}
	if voucherTemplateID != "" {
		filters["voucher_template_id"] = voucherTemplateID
	}

	if validFromStr != "" {
		if t, err := time.Parse("2006-01-02", validFromStr); err == nil {
			filters["valid_from"] = t
		}
	}

	if validUntilStr != "" {
		if t, err := time.Parse("2006-01-02", validUntilStr); err == nil {
			filters["valid_until"] = t
		}
	}

	vouchers, total, totalFiltered, err := c.service.FindAll(ctx.Context(), filters)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to fetch vouchers")
	}

	payload := fiber.Map{
		"vouchers":      vouchers,
		"total":         total,
		"totalFiltered": totalFiltered,
	}

	return helpers.WriteJSON(ctx, fiber.StatusOK, payload, "Vouchers fetched successfully")

}

func (c *VoucherControllerImpl) FindByID(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid voucher ID")
	}
	voucher, err := c.service.FindByID(ctx.Context(), id)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusNotFound, "Voucher not found")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, voucher, "Voucher fetched successfully")
}

func (c *VoucherControllerImpl) FindByCode(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	voucher, err := c.service.FindByCode(ctx.Context(), code)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusNotFound, "Voucher not found")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, voucher, "Voucher fetched successfully")
}

func (c *VoucherControllerImpl) Create(ctx *fiber.Ctx) error {
	var voucher domain.Voucher
	if err := ctx.BodyParser(&voucher); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request payload")
	}
	// voucher.Code = helpers.GenerateVoucherCode() // opsional, jika ingin auto-generate
	createdVoucher, err := c.service.Create(ctx.Context(), voucher)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to create voucher")
	}
	return helpers.WriteJSON(ctx, fiber.StatusCreated, createdVoucher, "Voucher created successfully")
}

func (c *VoucherControllerImpl) Update(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid voucher ID")
	}
	var voucher domain.Voucher
	if err := ctx.BodyParser(&voucher); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request payload")
	}
	voucher.ID = id
	updatedVoucher, err := c.service.Update(ctx.Context(), voucher)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to update voucher")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, updatedVoucher, "Voucher updated successfully")
}

func (c *VoucherControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid voucher ID")
	}
	if err := c.service.Delete(ctx.Context(), id); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to delete voucher")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, nil, "Voucher deleted successfully")
}

func (c *VoucherControllerImpl) UseVoucher(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid voucher ID")
	}

	type Request struct {
		BoothID string `json:"booth_id"`
	}

	var req Request
	if err := ctx.BodyParser(&req); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request payload")
	}

	updatedVoucher, err := c.service.UseVoucher(ctx.Context(), id, req.BoothID)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, err.Error())
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, updatedVoucher, "Voucher usage updated")
}
