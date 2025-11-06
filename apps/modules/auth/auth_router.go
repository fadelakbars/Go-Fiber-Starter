package auth

import (
	"go-fiber-starter/apps/middleware"
	"go-fiber-starter/apps/modules/auth/controller"
	"go-fiber-starter/apps/modules/auth/service"
	"go-fiber-starter/apps/modules/user/repository"

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
