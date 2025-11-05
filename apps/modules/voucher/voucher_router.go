package voucher

import (
	"mou-be/apps/middleware"
	"mou-be/apps/modules/voucher/controller"
	"mou-be/apps/modules/voucher/repository"
	"mou-be/apps/modules/voucher/service"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Router(app fiber.Router, db *gorm.DB, rdb *redis.Client) {
	repo := repository.NewVoucherRepository()
	svc := service.NewVoucherService(repo, db)
	ctrl := controller.NewVoucherController(svc)

	r := app.Group("/vouchers", middleware.APIKeyMiddleware())

	r.Get("/", ctrl.FindAll)
	r.Get("/:id", ctrl.FindByID)
	r.Get("/code/:code", ctrl.FindByCode)
	r.Post("/", ctrl.Create)
	r.Put("/:id", ctrl.Update)
	r.Put("/:id/use", ctrl.UseVoucher)
	r.Delete("/:id", ctrl.Delete)
}
