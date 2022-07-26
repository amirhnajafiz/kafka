package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) UpsertProject(c *fiber.Ctx) error {
	err := h.ProjectsCollection.Upsert(c.Context(), nil)
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.SendStatus(http.StatusOK)
}

func (h *Handler) RemoveProject(c *fiber.Ctx) error {
	err := h.ProjectsCollection.Delete(c.Context(), "")
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.SendStatus(http.StatusAccepted)
}
