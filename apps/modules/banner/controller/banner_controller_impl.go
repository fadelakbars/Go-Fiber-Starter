package controller

import (
	"mou-be/apps/domain"
	"mou-be/apps/helpers"
	"mou-be/apps/modules/banner/service"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type BannerControllerImpl struct {
	service    service.BannerService
	s3Client   *s3.S3
	bucketName string
}

func NewBannerController(service service.BannerService, s3Client *s3.S3, bucketName string) BannerController {
	return &BannerControllerImpl{
		service:    service,
		s3Client:   s3Client,
		bucketName: bucketName,
	}
}

func (c *BannerControllerImpl) FindAll(ctx *fiber.Ctx) error {
	banner, err := c.service.FindAll(ctx.Context())
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to fetch banner")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, banner, "Banner fetched successfully")
}

func (c *BannerControllerImpl) Create(ctx *fiber.Ctx) error {
	name := ctx.FormValue("name")

	// Handle file upload
	file, err := ctx.FormFile("image")
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Image file is required")
	}
	imageUrl, err := helpers.UploadFileToS3(c.s3Client, c.bucketName, file, "banner")
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to upload image")
	}

	banner := domain.Banner{
		Name:     name,
		ImageUrl: imageUrl,
	}

	createdBanner, err := c.service.Create(ctx.Context(), banner)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to create banner")
	}

	return helpers.WriteJSON(ctx, fiber.StatusCreated, createdBanner, "Banner created successfully")
}

func (c *BannerControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid banner ID")
	}

	if err := c.service.Delete(ctx.Context(), id); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to delete banner")
	}

	return helpers.WriteJSON(ctx, fiber.StatusOK, nil, "Banner deleted successfully")
}
