package router

import (
	auth "mou-be/apps/modules/auth"
	banner "mou-be/apps/modules/banner"
	booth "mou-be/apps/modules/booth"
	order "mou-be/apps/modules/order"
	payment "mou-be/apps/modules/payment"
	test "mou-be/apps/modules/test"
	user "mou-be/apps/modules/user"
	voucher "mou-be/apps/modules/voucher"
	voucher_template "mou-be/apps/modules/voucher_template"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB, rdb *redis.Client, s3Client *s3.S3, bucketName string) {

	apiRoutes := app.Group("/api/v1")

	user.Router(apiRoutes, db, rdb)
	booth.Router(apiRoutes, db, rdb, s3Client, bucketName)
	auth.Router(apiRoutes, db, rdb)
	payment.Router(apiRoutes, db, rdb)
	voucher.Router(apiRoutes, db, rdb)
	banner.Router(apiRoutes, db, rdb, s3Client, bucketName)
	order.Router(apiRoutes, db, rdb)
	voucher_template.Router(apiRoutes, db, rdb, s3Client, bucketName)

	test.Router(apiRoutes, db, rdb)

}
