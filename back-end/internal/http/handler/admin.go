package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Login(c *fiber.Ctx) error {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var userRequest request

	_ = c.BodyParser(&userRequest)

	if userRequest.Username != h.Cfg.Admin.User || userRequest.Password != h.Cfg.Admin.Pass {
		return c.SendStatus(http.StatusUnauthorized)
	}

	token, err := h.JWT.GenerateToken(userRequest.Username, userRequest.Password)
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.SendString(token)
}
