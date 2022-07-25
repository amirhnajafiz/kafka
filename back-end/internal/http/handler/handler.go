package handler

import (
	"github.com/gofiber/fiber/v2"
)

type Handler struct{}

func (h *Handler) RegisterPersonal(api fiber.Router) {
	api.Get("/personal", h.GetPersonal)
	api.Get("/projects", h.GetProjects)
	api.Get("/project/:id", h.GetProjectById)
}

func (h *Handler) RegisterAdmin(api fiber.Router) {
	api.Put("/personal", h.UpsertProject)
	api.Put("/project/:id", h.RemoveProject)
	api.Delete("/project/:id", h.RemoveProject)
}
