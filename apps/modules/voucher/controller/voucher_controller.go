package controller

import "github.com/gofiber/fiber/v2"

type VoucherController interface {
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	FindByCode(ctx *fiber.Ctx) error
	UseVoucher(ctx *fiber.Ctx) error
}
