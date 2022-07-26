package cmd

import (
	"log"

	"github.com/amirhnajafiz/personal-website/back-end/internal/cmd/server"
	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "back-end",
		Short: "personal website backend",
		Long:  "starting the personal website backend",
	}

	rootCmd.AddCommand(server.Command())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
