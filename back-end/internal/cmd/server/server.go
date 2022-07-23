package server

import (
	"github.com/amirhnajafiz/personal-website/back-end/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	app := fiber.New()

	h := handler.Handler{}

	app.Get("/api/personalInfo", h.GetPersonalInformation)
	app.Get("/api/projects", h.GetListOfProjects)
	app.Get("/api/projects/:id", h.GetASingleProject)

	return app
}
