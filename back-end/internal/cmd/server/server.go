package server

import (
	"github.com/amirhnajafiz/personal-website/back-end/internal/http/handler"
	"github.com/amirhnajafiz/personal-website/back-end/internal/http/middleware"
	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	// create app
	app := fiber.New()

	// using auth middleware
	app.Use("/api/admin", middleware.Authentication)

	// creating our api and handler
	api := app.Group("/api")
	h := handler.Handler{}

	// register our handler
	h.Register(api)

	return app
}
