package migrate

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

const (
	path = "file://./internal/model/migrations"
	url  = ""
)

// Command will return the cobra command for migration package
func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "migrate mongodb",
		Long:  "migrate our website mongo database",
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
	m, err := migrate.New(
		path,
		url,
	)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	var migration func() error

	if args[0] == "up" {
		migration = m.Up
	} else {
		migration = m.Down
	}

	if err := migration(); err != nil {
		return fmt.Errorf("failed to execute migration: %w", err)
	}

	return nil
}
