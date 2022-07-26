package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "serve",
		Short: "personal website backend",
		Long:  "starting the personal website backend server",
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
