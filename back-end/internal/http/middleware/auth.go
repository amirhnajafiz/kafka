package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Authentication(c *fiber.Ctx) error {
	return c.Next()
}
