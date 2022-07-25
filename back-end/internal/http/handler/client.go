package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetProjects(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func (h *Handler) GetProjectById(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusAccepted)
}
