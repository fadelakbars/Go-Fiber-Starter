package helpers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleError(ctx *fiber.Ctx, err error, statusCode int, message string) error {
	if err != nil {
		log.Println(err)
	}
	return WriteJSON(ctx, statusCode, nil, message)
}
