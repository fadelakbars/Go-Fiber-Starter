package user

import (
	"mou-be/apps/middleware"
	"mou-be/apps/modules/user/controller"
	"mou-be/apps/modules/user/repository"
	"mou-be/apps/modules/user/service"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Router(app fiber.Router, db *gorm.DB, rdb *redis.Client) {
	repo := repository.NewUserRepository()
	svc := service.NewUserService(repo, db)
	ctrl := controller.NewUserController(svc)

	r := app.Group("/users", middleware.APIKeyMiddleware())

	r.Get("/", ctrl.FindAll)
	r.Get("/:id", ctrl.FindByID)
	r.Post("/", ctrl.Create)
	r.Put("/:id", ctrl.Update)
	r.Delete("/:id", ctrl.Delete)
}
