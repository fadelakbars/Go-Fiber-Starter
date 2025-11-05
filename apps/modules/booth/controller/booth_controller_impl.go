package controller

import (
	"encoding/json"
	"mou-be/apps/domain"
	"mou-be/apps/helpers"
	"mou-be/apps/modules/booth/service"
	"strconv"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type BoothControllerImpl struct {
	service    service.BoothService
	s3Client   *s3.S3
	bucketName string
}

func NewBoothController(service service.BoothService, s3Client *s3.S3, bucketName string) BoothController {
	return &BoothControllerImpl{
		service:    service,
		s3Client:   s3Client,
		bucketName: bucketName,
	}
}

func (c *BoothControllerImpl) FindAll(ctx *fiber.Ctx) error {
	booths, err := c.service.FindAll(ctx.Context())
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to fetch booths")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, booths, "Booths fetched successfully")
}

func (c *BoothControllerImpl) FindByID(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid booth ID")
	}
	booth, err := c.service.FindByID(ctx.Context(), id)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusNotFound, "Booth not found")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, booth, "Booth fetched successfully")
}

func (c *BoothControllerImpl) Create(ctx *fiber.Ctx) error {
	var booth domain.Booth

	// Parse form fields
	booth.BoothName = ctx.FormValue("booth_name")
	booth.BoothUsername = ctx.FormValue("booth_username")
	booth.BoothDeviceID = ctx.FormValue("booth_device_id")
	booth.PIN = ctx.FormValue("pin")
	booth.Location = ctx.FormValue("location")
	booth.PhoneNumber = ctx.FormValue("phone_number")
	booth.CameraScanner = ctx.FormValue("camera_scanner")
	booth.CampaignUrl = ctx.FormValue("campaign_url")
	// Parse float (price)
	priceStr := ctx.FormValue("price")
	if priceStr != "" {
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid price format")
		}
		booth.Price = price
	}

	// Parse int (timeout fields)
	paymentTimeout := ctx.FormValue("payment_timeout_seconds")
	if paymentTimeout != "" {
		val, err := strconv.Atoi(paymentTimeout)
		if err != nil {
			return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid payment_timeout_seconds")
		}
		booth.PaymentTimeoutSeconds = val
	}

	dslrTimeout := ctx.FormValue("dslrbooth_timeout_seconds")
	if dslrTimeout != "" {
		val, err := strconv.Atoi(dslrTimeout)
		if err != nil {
			return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid dslrbooth_timeout_seconds")
		}
		booth.DslrBoothTimeoutSeconds = val
	}

	primaryColorStr := ctx.FormValue("primary_color")
	if primaryColorStr != "" {
		var js json.RawMessage
		if err := json.Unmarshal([]byte(primaryColorStr), &js); err != nil {
			return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid JSON for primary_color")
		}
		booth.PrimaryColor = datatypes.JSON(js)
	}

	// Upload images if present
	imageStartFile, _ := ctx.FormFile("image_start")
	if imageStartFile != nil {
		url, err := helpers.UploadFileToS3(c.s3Client, c.bucketName, imageStartFile, "booth")
		if err != nil {
			return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to upload image_start")
		}
		booth.ImageStart = url
	}
	imageContentFile, _ := ctx.FormFile("image_content")
	if imageContentFile != nil {
		url, err := helpers.UploadFileToS3(c.s3Client, c.bucketName, imageContentFile, "booth")
		if err != nil {
			return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to upload image_content")
		}
		booth.ImageContent = url
	}
	imageFooterFile, _ := ctx.FormFile("image_footer")
	if imageFooterFile != nil {
		url, err := helpers.UploadFileToS3(c.s3Client, c.bucketName, imageFooterFile, "booth")
		if err != nil {
			return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to upload image_footer")
		}
		booth.Imagefooter = url
	}

	createdBooth, err := c.service.Create(ctx.Context(), booth)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to create booth")
	}
	return helpers.WriteJSON(ctx, fiber.StatusCreated, createdBooth, "Booth created successfully")
}

func (c *BoothControllerImpl) Update(ctx *fiber.Ctx) error {
	// Parse ID dari parameter URL
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid booth ID")
	}

	// Ambil data lama dari service
	oldBooth, err := c.service.FindByID(ctx.Context(), id)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusNotFound, "Booth not found")
	}

	// Update nilai berdasarkan form
	booth := oldBooth // start dari data lama

	// Text fields
	if v := ctx.FormValue("booth_name"); v != "" {
		booth.BoothName = v
	}
	if v := ctx.FormValue("booth_username"); v != "" {
		booth.BoothUsername = v
	}
	if v := ctx.FormValue("booth_device_id"); v != "" {
		booth.BoothDeviceID = v
	}
	if v := ctx.FormValue("pin"); v != "" {
		booth.PIN = v
	}
	if v := ctx.FormValue("location"); v != "" {
		booth.Location = v
	}
	if v := ctx.FormValue("phone_number"); v != "" {
		booth.PhoneNumber = v
	}
	if v := ctx.FormValue("camera_scanner"); v != "" {
		booth.CameraScanner = v
	}
	if v := ctx.FormValue("campaign_url"); v != "" {
		booth.CampaignUrl = v
	}

	// Numeric fields
	if priceStr := ctx.FormValue("price"); priceStr != "" {
		if price, err := strconv.ParseFloat(priceStr, 64); err == nil {
			booth.Price = price
		} else {
			return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid price format")
		}
	}
	if timeout := ctx.FormValue("payment_timeout_seconds"); timeout != "" {
		if val, err := strconv.Atoi(timeout); err == nil {
			booth.PaymentTimeoutSeconds = val
		} else {
			return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid payment_timeout_seconds")
		}
	}
	if timeout := ctx.FormValue("dslrbooth_timeout_seconds"); timeout != "" {
		if val, err := strconv.Atoi(timeout); err == nil {
			booth.DslrBoothTimeoutSeconds = val
		} else {
			return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid dslrbooth_timeout_seconds")
		}
	}

	// JSON field
	if primaryColorStr := ctx.FormValue("primary_color"); primaryColorStr != "" {
		var js json.RawMessage
		if err := json.Unmarshal([]byte(primaryColorStr), &js); err != nil {
			return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid JSON for primary_color")
		}
		booth.PrimaryColor = datatypes.JSON(js)
	}

	// Handle image uploads
	uploadImage := func(fieldName, oldURL string) (string, error) {
		file, err := ctx.FormFile(fieldName)
		if err != nil {
			return oldURL, nil // skip if no new file
		}
		if oldURL != "" {
			_ = helpers.DeleteFileFromS3(c.s3Client, c.bucketName, oldURL) // optional: ignore error
		}
		newURL, err := helpers.UploadFileToS3(c.s3Client, c.bucketName, file, "booth")
		if err != nil {
			return "", err
		}
		return newURL, nil
	}

	if booth.ImageStart, err = uploadImage("image_start", booth.ImageStart); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to upload image_start")
	}
	if booth.ImageContent, err = uploadImage("image_content", booth.ImageContent); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to upload image_content")
	}
	if booth.Imagefooter, err = uploadImage("image_footer", booth.Imagefooter); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to upload image_footer")
	}

	// Save updated booth
	updatedBooth, err := c.service.Update(ctx.Context(), booth)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to update booth")
	}

	return helpers.WriteJSON(ctx, fiber.StatusOK, updatedBooth, "Booth updated successfully")
}

// func (c *BoothControllerImpl) Update(ctx *fiber.Ctx) error {
// 	id, err := uuid.Parse(ctx.Params("id"))
// 	if err != nil {
// 		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid booth ID")
// 	}

// 	var booth domain.Booth
// 	if err := ctx.BodyParser(&booth); err != nil {
// 		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request payload")
// 	}

// 	// Ambil body mentah sebagai map
// 	var body map[string]interface{}
// 	if err := ctx.BodyParser(&body); err == nil {
// 		if pc, ok := body["primary_color"]; ok {
// 			jsonBytes, err := json.Marshal(pc)
// 			if err != nil {
// 				return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid primary_color format")
// 			}
// 			booth.PrimaryColor = datatypes.JSON(jsonBytes)
// 		}
// 	}

// 	booth.ID = id
// 	updatedBooth, err := c.service.Update(ctx.Context(), booth)
// 	if err != nil {
// 		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to update booth")
// 	}
// 	return helpers.WriteJSON(ctx, fiber.StatusOK, updatedBooth, "Booth updated successfully")
// }

func (c *BoothControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid booth ID")
	}
	if err := c.service.Delete(ctx.Context(), id); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to delete booth")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, nil, "Booth deleted successfully")
}

func (c *BoothControllerImpl) Login(ctx *fiber.Ctx) error {
	var req struct {
		BoothUsername string `json:"booth_username"`
		PIN           string `json:"pin"`
		BoothDeviceID string `json:"booth_device_id"`
	}

	if err := ctx.BodyParser(&req); err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request payload")
	}
	token, booth, err := c.service.Login(ctx.Context(), req.BoothUsername, req.PIN, req.BoothDeviceID)
	if err != nil {
		return helpers.HandleError(ctx, err, fiber.StatusUnauthorized, "Invalid booth username or PIN")
	}
	return helpers.WriteJSON(ctx, fiber.StatusOK, fiber.Map{
		"token": token,
		"booth": booth,
	}, "Login successful")
}
