package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// APIKeyMiddleware middleware untuk validasi X-API-Key header
func APIKeyMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		apiKey := c.Get("X-API-Key")

		if apiKey == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing API Key",
			})
		}

		// Ganti sesuai kebutuhan atau load dari env/config file
		validKey := "Go-Fiber-Starter-API-Key"

		if apiKey != validKey {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid API Key",
			})
		}

		return c.Next()
	}
}
