package handler

import (
	"net/http"

	"github.com/amirhnajafiz/personal-website/back-end/internal/model"
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
	type request struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Link        string `json:"link"`
	}

	var userRequest request

	if err := c.BodyParser(&userRequest); err != nil {
		return c.SendString(err.Error())
	}

	p := &model.Project{
		Title:       userRequest.Title,
		Description: userRequest.Description,
		Link:        userRequest.Link,
		Visible:     false,
	}

	err := h.ProjectsCollection.Upsert(c.Context(), p)
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.SendStatus(http.StatusOK)
}

func (h *Handler) RemoveProject(c *fiber.Ctx) error {
	type request struct {
		Title string `json:"title"`
	}

	var userRequest request

	if err := c.BodyParser(&userRequest); err != nil {
		return c.SendString(err.Error())
	}

	err := h.ProjectsCollection.Delete(c.Context(), userRequest.Title)
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.SendStatus(http.StatusAccepted)
}
