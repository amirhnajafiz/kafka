package handler

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetProjects(c *fiber.Ctx) error {
	p, err := h.ProjectsCollection.GetAllAvailable(c.Context())
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.JSON(p)
}

func (h *Handler) GetProjectById(c *fiber.Ctx) error {
	p, err := h.ProjectsCollection.GetSingle(c.Context(), "")
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.JSON(p)
}
