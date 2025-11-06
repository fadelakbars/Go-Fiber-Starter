package main

import (
	"fmt"
	"go-fiber-starter/apps/config"
	"go-fiber-starter/apps/middleware"
	"go-fiber-starter/apps/router"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
}

func setup() *fiber.App {
	db := config.DBConnect()
	middleware.InitSecretKey()

	app := fiber.New()

	// Tambahkan middleware CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8080,https://panel-mou-dev.teknozen.id", // Bisa diganti dengan domain tertentu, misal "https://example.com"
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Authorization,X-API-Key",
		AllowCredentials: true, // Jika perlu mengirim credentials seperti cookies atau Authorization header
	}))

	// Setup routes
	router.SetupRoutes(app, db)

	return app
}

func main() {
	app := setup()

	app_port, _ := strconv.Atoi(viper.Get("APP_PORT").(string))

	log.Fatal(app.Listen(fmt.Sprintf(":%d", app_port)))
}
