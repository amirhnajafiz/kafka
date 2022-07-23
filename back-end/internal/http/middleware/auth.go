package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Authentication(c *fiber.Ctx) error {
	if _, ok := c.GetReqHeaders()["jwt"]; ok {
		return c.Next()
	}

	return c.SendStatus(http.StatusUnauthorized)
}
