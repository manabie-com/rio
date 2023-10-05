package test

import (
	"context"
	"path/filepath"

	"github.com/manabie-com/rio/internal/config"
	"github.com/manabie-com/rio/internal/database"
)

// ResetDB resets DB for testing
func ResetDB(ctx context.Context, basePath string) {
	dbConfig := config.NewDBConfig()
	gormDB, err := database.Connect(ctx, dbConfig)
	if err != nil {
		panic(err)
	}

	err = database.ExecuteFileScript(ctx, gormDB, filepath.Join(basePath, "schema/reset_db.sql"))
	if err != nil {
		panic(err)
	}

	if err := database.Migrate(ctx, dbConfig, filepath.Join(basePath, "schema/migration")); err != nil {
		panic(err)
	}
}
