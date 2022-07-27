package migrate

import (
	"fmt"
	"log"

	"github.com/amirhnajafiz/personal-website/back-end/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

const (
	path = "file://./internal/model/migrations"
)

// Command will return the cobra command for migration package
func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "migrate mongodb",
		Long:  "executing migration for our mongodb",
		Run: func(cmd *cobra.Command, args []string) {
			err := run(args)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
}

// run is the entrypoint of our migrate package
func run(args []string) error {
	// loading configs
	cfg := config.Load()

	// creating a new migrate
	m, err := migrate.New(
		path,
		cfg.Mongodb.URL,
	)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	// migration function
	var migration func() error

	if args[0] == "up" {
		migration = m.Up
	} else {
		migration = m.Down
	}

	// executing migration
	if err := migration(); err != nil {
		return fmt.Errorf("failed to execute migration: %w", err)
	}

	return nil
}
