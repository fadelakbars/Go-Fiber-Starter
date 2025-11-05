package voucher_template

import (
	"mou-be/apps/middleware"
	"mou-be/apps/modules/voucher_template/controller"
	"mou-be/apps/modules/voucher_template/repository"
	"mou-be/apps/modules/voucher_template/service"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Router(app fiber.Router, db *gorm.DB, rdb *redis.Client, s3Client *s3.S3, bucketName string) {
	repo := repository.NewVoucherTemplateRepository()
	svc := service.NewVoucherTemplateService(repo, db)
	ctrl := controller.NewVoucherTemplateController(svc, s3Client, bucketName)

	r := app.Group("/voucher-templates", middleware.APIKeyMiddleware())
	r.Get("/", ctrl.FindAll)
	r.Get("/:id", ctrl.FindByID)
	r.Post("/", ctrl.Create)
	r.Put("/:id", ctrl.Update)
	r.Delete("/:id", ctrl.Delete)
}
