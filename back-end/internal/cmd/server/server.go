package server

import (
	"github.com/amirhnajafiz/personal-website/back-end/internal/http/handler"
	"github.com/amirhnajafiz/personal-website/back-end/internal/http/middleware"
	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	app := fiber.New()

	app.Use("/api/admin", middleware.Authentication)

	api := app.Group("/api")
	h := handler.Handler{}

	h.Register(api)

	return app
}
