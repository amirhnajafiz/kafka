package handler

import (
	"github.com/amirhnajafiz/personal-website/back-end/internal/database/store"
	"github.com/amirhnajafiz/personal-website/back-end/internal/jwt"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Cfg                Config
	JWT                jwt.JWT
	ProjectsCollection store.ProjectsCollection
}

// RegisterClient creates the routes of client
func (h *Handler) RegisterClient(api fiber.Router) {
	api.Get("/projects", h.GetVisibleProjects)
	api.Get("/project/:title", h.GetProjectById)
	api.Post("/login", h.Login)
}

// RegisterAdmin creates the routes of admin
func (h *Handler) RegisterAdmin(api fiber.Router) {
	api.Get("/projects", h.GetAllProjects)
	api.Put("/projects", h.UpsertProject)
	api.Delete("/project/:title", h.RemoveProject)
}
