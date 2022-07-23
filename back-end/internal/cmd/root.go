package cmd

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Execute() {
	app := fiber.New()

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
