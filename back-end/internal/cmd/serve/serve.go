package serve

import (
	"log"
	"strconv"

	"github.com/amirhnajafiz/personal-website/back-end/internal/http/handler"
	"github.com/amirhnajafiz/personal-website/back-end/internal/http/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

// Command will return the cobra command for serve package.
func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "start backend server",
		Long:  "start backend server",
		Run: func(_ *cobra.Command, _ []string) {
			run()
		},
	}
}

// entrypoint of the serve command
func run() {
	// create app
	app := fiber.New()

	// using auth middleware
	app.Use("/api/admin", middleware.Authentication)

	// creating our api and handler
	v1 := app.Group("/api")
	v2 := v1.Group("/admin")
	h := handler.Handler{}

	// register our handler
	h.RegisterClient(v1)
	h.RegisterAdmin(v2)

	serve(app, 8080)
}

// serve starts the backend server on a give port
func serve(app *fiber.App, port int) {
	addr := ":" + strconv.Itoa(port)

	if err := app.Listen(addr); err != nil {
		log.Fatal(err)
	}
}
