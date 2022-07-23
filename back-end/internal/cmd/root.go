package cmd

import (
	"log"

	"github.com/amirhnajafiz/personal-website/back-end/internal/cmd/server"
)

func Execute() {
	app := server.New()

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
