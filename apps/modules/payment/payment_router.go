package payment

import (
	"mou-be/apps/middleware"
	"mou-be/apps/modules/payment/controller"
	"mou-be/apps/modules/payment/repository"
	"mou-be/apps/modules/payment/service"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Router(app fiber.Router, db *gorm.DB, rdb *redis.Client) {
	repo := repository.NewPaymentRepository()
	svc := service.NewPaymentService(repo, db)
	ctrl := controller.NewPaymentController(svc)

	r := app.Group("/payments")

	r.Get("/", ctrl.FindAll)
	r.Get("/:id", ctrl.FindByID)
	// r.Post("/", ctrl.Create)
	// r.Put("/:id", ctrl.Update)
	// r.Delete("/:id", ctrl.Delete)
	r.Post("/snap", ctrl.CreateSnapPayment, middleware.APIKeyMiddleware())
	r.Post("/callback", ctrl.HandleMidtransCallback)

	// test
}
