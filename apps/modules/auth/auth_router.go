package auth

import (
	"mou-be/apps/middleware"
	"mou-be/apps/modules/auth/controller"
	"mou-be/apps/modules/auth/service"
	"mou-be/apps/modules/user/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(app fiber.Router, db *gorm.DB) {
	repo := repository.NewUserRepository()
	svc := service.NewAuthService(repo, db)
	ctrl := controller.NewAuthController(svc)

	r := app.Group("/auth", middleware.APIKeyMiddleware())
	r.Post("/login", ctrl.Login)
}
