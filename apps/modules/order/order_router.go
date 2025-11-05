package order

import (
	"mou-be/apps/middleware"
	"mou-be/apps/modules/order/controller"
	"mou-be/apps/modules/order/repository"
	"mou-be/apps/modules/order/service"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Router(app fiber.Router, db *gorm.DB, rdb *redis.Client) {
	repo := repository.NewOrderRepository()
	svc := service.NewOrderService(repo, db)
	ctrl := controller.NewOrderController(svc)

	r := app.Group("/orders", middleware.APIKeyMiddleware())

	r.Get("/", ctrl.FindAll)
	// r.Get("/:id", ctrl.FindByID)
	// r.Post("/", ctrl.Create)
	// r.Put("/:id", ctrl.Update)
	// r.Delete("/:id", ctrl.Delete)
}
