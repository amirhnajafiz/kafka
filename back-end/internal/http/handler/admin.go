package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *Handler) GetAllProjects(c *fiber.Ctx) error {
	filter := bson.D{{}}

	projects, err := h.ProjectsCollection.GetAll(c.Context(), filter)
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.JSON(projects)
}

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
