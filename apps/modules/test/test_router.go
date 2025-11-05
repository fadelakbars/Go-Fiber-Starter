package test

import (
	"mou-be/apps/middleware"
	"mou-be/apps/modules/test/controller"
	"mou-be/apps/modules/test/repository"
	"mou-be/apps/modules/test/service"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Router(app fiber.Router, db *gorm.DB, rdb *redis.Client) {
	repo := repository.NewTestRepository()
	svc := service.NewTestService(repo, db)
	ctrl := controller.NewTestController(svc)

	r := app.Group("/tests", middleware.APIKeyMiddleware())

	r.Get("/", ctrl.FindAll)
	r.Get("/:id", ctrl.FindByID)
	r.Post("/", ctrl.Create)
	r.Put("/:id", ctrl.Update)
	r.Delete("/:id", ctrl.Delete)
}
