package migrate

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	path = "file://./internal/model/migrations"
)

func Migrate(flag bool, url string) error {
	m, err := migrate.New(
		path,
		url,
	)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	var migration func() error

	if flag {
		migration = m.Up
	} else {
		migration = m.Down
	}

	if err := migration(); err != nil {
		return fmt.Errorf("failed to execute migration: %w", err)
	}

	return nil
}
