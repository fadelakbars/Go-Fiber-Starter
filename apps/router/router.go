package router

import (
	auth "mou-be/apps/modules/auth"
	user "mou-be/apps/modules/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {

	apiRoutes := app.Group("/api/v1")

	user.Router(apiRoutes, db)
	auth.Router(apiRoutes, db)

}
