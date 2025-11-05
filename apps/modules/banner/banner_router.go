package subcategory

import (
	"mou-be/apps/middleware"
	"mou-be/apps/modules/banner/controller"
	"mou-be/apps/modules/banner/repository"
	"mou-be/apps/modules/banner/service"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Router(app fiber.Router, db *gorm.DB, rdb *redis.Client, s3Client *s3.S3, bucketName string) {
	repo := repository.NewBannerRepository()
	svc := service.NewBannerService(repo, db)
	ctrl := controller.NewBannerController(svc, s3Client, bucketName)

	r := app.Group("/banner", middleware.APIKeyMiddleware())
	r.Get("/", ctrl.FindAll)
	r.Post("/", ctrl.Create)
	r.Delete("/:id", ctrl.Delete)
}
