package cmd

import (
	"log"

	"github.com/amirhnajafiz/personal-website/back-end/internal/cmd/migrate"
	"github.com/amirhnajafiz/personal-website/back-end/internal/cmd/serve"
	"github.com/spf13/cobra"
)

func Execute() {
	// creating a root command for cobra
	rootCmd := &cobra.Command{
		Use:   "personal website back-end",
		Short: "personal website back-end",
		Long:  "starting the personal website back-end",
	}

	// adding the serve and migrate commands
	rootCmd.AddCommand(
		serve.Command(),
		migrate.Command(),
	)

	// executing command
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
