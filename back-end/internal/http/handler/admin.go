package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) UpdatePersonal(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func (h *Handler) UpdateProject(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func (h *Handler) RemoveProject(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusAccepted)
}
