package router

import (
	auth "go-fiber-starter/apps/modules/auth"
	user "go-fiber-starter/apps/modules/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {

	apiRoutes := app.Group("/api/v1")

	user.Router(apiRoutes, db)
	auth.Router(apiRoutes, db)

}
