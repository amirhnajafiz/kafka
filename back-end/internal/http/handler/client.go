package handler

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *Handler) GetVisibleProjects(c *fiber.Ctx) error {
	filter := bson.M{
		"visible": bson.M{
			"$eq": true,
		},
	}

	p, err := h.ProjectsCollection.GetAll(c.Context(), filter)
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.JSON(p)
}

func (h *Handler) GetProjectById(c *fiber.Ctx) error {
	p, err := h.ProjectsCollection.GetSingle(c.Context(), c.Params("title"))
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.JSON(p)
}
