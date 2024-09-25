package middlewares

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("Request URL:", c.OriginalURL())
		return c.Next()
	}
}
