package user

import (
	"go-fiber-starter/apps/middleware"
	"go-fiber-starter/apps/modules/user/controller"
	"go-fiber-starter/apps/modules/user/repository"
	"go-fiber-starter/apps/modules/user/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(app fiber.Router, db *gorm.DB) {
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
