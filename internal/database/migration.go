package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/manabie-com/rio/internal/config"
	"github.com/manabie-com/rio/internal/log"
	"net/url"

	// blank import
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(ctx context.Context, config *config.PostgresConfig, dir string) error {
	connectionString := resolveDatabaseConnectionURL(config)
	m, err := migrate.New(fmt.Sprintf("file://%s", dir), connectionString)
	if err != nil {
		return err
	}

	defer m.Close()

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Info(ctx, "no migration needed")
			return nil
		}

		return err
	}

	return nil
}

func resolveDatabaseConnectionURL(config *config.PostgresConfig) string {
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable&application_name=%s",
		url.QueryEscape(config.User),
		url.QueryEscape(config.Password),
		config.Host,
		config.Port,
		url.QueryEscape(config.Schema),
		url.QueryEscape(config.User),
	)
	return dsn
}
