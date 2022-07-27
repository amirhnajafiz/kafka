package middleware

import (
	"net/http"

	"github.com/amirhnajafiz/personal-website/back-end/internal/jwt"
	"github.com/gofiber/fiber/v2"
)

type Auth struct {
	JWT jwt.JWT
}

func (a *Auth) Authentication(c *fiber.Ctx) error {
	if token, ok := c.GetReqHeaders()["jwt"]; ok {
		if valid, _ := a.JWT.ParseToken(token); valid {
			return c.Next()
		}
	}

	return c.SendStatus(http.StatusUnauthorized)
}
