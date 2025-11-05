package middleware

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var SecretKey string

// InitSecretKey initializes the SecretKey from the configuration.
func InitSecretKey() {
	SecretKey = viper.GetString("JWT_SECRET_KEY")
	if SecretKey == "" {
		log.Fatal("JWT_SECRET_KEY is not set in the configuration")
	}
}

// AuthMiddleware is a middleware for authenticating requests using JWT and Redis.
func AuthMiddleware(rdb *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString, err := getTokenFromHeader(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		claims, err := parseAndValidateToken(tokenString)
		if err != nil {
			log.Println("Token validation error:", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		userID, ok := claims["userID"].(string)
		if !ok {
			log.Println("Invalid userID claim in token")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		log.Println("UserID from token:", userID)

		// err = checkTokenInRedis(rdb, userID, tokenString)
		// if err != nil {
		// 	log.Println("Token not found in Redis or does not match:", err)
		// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		// }

		c.Locals("userID", userID)
		return c.Next()
	}
}

// getTokenFromHeader retrieves the JWT from the Authorization header.
func getTokenFromHeader(c *fiber.Ctx) (string, error) {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}
	if len(tokenString) <= len("Bearer ") {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Invalid token format")
	}
	return tokenString[len("Bearer "):], nil
}

// parseAndValidateToken parses and validates the JWT.
func parseAndValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	return claims, nil
}

// checkTokenInRedis checks if the token is stored in Redis.
func checkTokenInRedis(rdb *redis.Client, userID string, tokenString string) error {
	storedToken, err := rdb.Get(context.Background(), userID).Result()
	if err == redis.Nil || storedToken != tokenString {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	} else if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}
	return nil
}
