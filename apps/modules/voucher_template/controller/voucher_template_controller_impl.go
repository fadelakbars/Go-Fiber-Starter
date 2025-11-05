package controller

import (
	"mou-be/apps/domain"
	"mou-be/apps/helpers"
	"mou-be/apps/modules/voucher_template/service"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type VoucherTemplateControllerImpl struct {
	service    service.VoucherTemplateService
	s3Client   *s3.S3
	bucketName string
}

func NewVoucherTemplateController(service service.VoucherTemplateService, s3Client *s3.S3, bucketName string) VoucherTemplateController {
	return &VoucherTemplateControllerImpl{
		service:    service,
		s3Client:   s3Client,
		bucketName: bucketName,
	}
}

func (c *VoucherTemplateControllerImpl) FindAll(ctx *fiber.Ctx) error {
	items, err := c.service.FindAll(ctx.Context())
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to fetch voucher templates")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, items, "Voucher templates fetched successfully")
}

func (c *VoucherTemplateControllerImpl) FindByID(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid voucher template ID")
	}
	item, err := c.service.FindByID(ctx.Context(), id)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusNotFound, "Voucher template not found")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, item, "Voucher template fetched successfully")
}

func (c *VoucherTemplateControllerImpl) Create(ctx *fiber.Ctx) error {
	name := ctx.FormValue("name")
	horizontal := ctx.FormValue("horizontal")
	vertical := ctx.FormValue("vertical")
	size := ctx.FormValue("size")

	file, err := ctx.FormFile("image")
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Image file is required")
	}
	imageUrl, err := helpers.UploadFileToS3(c.s3Client, c.bucketName, file, "voucher-template")
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to upload image")
	}

	item := domain.VoucherTemplate{
		Name:       name,
		ImageUrl:   imageUrl,
		Horizontal: horizontal,
		Vertical:   vertical,
		Size:       size,
	}

	created, err := c.service.Create(ctx.Context(), item)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to create voucher template")
	}

	return helpers.WriteJSON(ctx, fiber.StatusCreated, created, "Voucher template created successfully")
}

func (c *VoucherTemplateControllerImpl) Update(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid voucher template ID")
	}

	name := ctx.FormValue("name")
	horizontal := ctx.FormValue("horizontal")
	vertical := ctx.FormValue("vertical")
	size := ctx.FormValue("size")

	var imageUrl string
	file, errFile := ctx.FormFile("image")
	if errFile == nil && file != nil {
		imageUrl, err = helpers.UploadFileToS3(c.s3Client, c.bucketName, file, "voucher-template")
		if err != nil {
			return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to upload image")
		}
	}

	item := domain.VoucherTemplate{
		ID:         id,
		Name:       name,
		ImageUrl:   imageUrl,
		Horizontal: horizontal,
		Vertical:   vertical,
		Size:       size,
	}

	updated, err := c.service.Update(ctx.Context(), item)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to update voucher template")
	}

	return helpers.WriteJSON(ctx, fiber.StatusOK, updated, "Voucher template updated successfully")
}

func (c *VoucherTemplateControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid voucher template ID")
	}

	if err := c.service.Delete(ctx.Context(), id); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to delete voucher template")
	}

	return helpers.WriteJSON(ctx, fiber.StatusOK, nil, "Voucher template deleted successfully")
}
