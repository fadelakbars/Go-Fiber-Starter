package booth

import (
	"mou-be/apps/middleware"
	"mou-be/apps/modules/booth/controller"
	"mou-be/apps/modules/booth/repository"
	"mou-be/apps/modules/booth/service"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// Tambahkan parameter s3Client dan bucketName pada Router
func Router(app fiber.Router, db *gorm.DB, rdb *redis.Client, s3Client *s3.S3, bucketName string) {
	repo := repository.NewBoothRepository()
	svc := service.NewBoothService(repo, db)
	ctrl := controller.NewBoothController(svc, s3Client, bucketName)

	r := app.Group("/booths", middleware.APIKeyMiddleware())

	r.Get("/", ctrl.FindAll)
	r.Get("/:id", ctrl.FindByID)
	r.Post("/", ctrl.Create)
	r.Post("/:id", ctrl.Update)
	r.Delete("/:id", ctrl.Delete)
	r.Post("/auth/login", ctrl.Login) // Tambahkan endpoint login booth
}
