package setup

import (
	"context"

	"github.com/manabie-com/rio"
	"github.com/manabie-com/rio/internal/cache"
	"github.com/manabie-com/rio/internal/config"
	"github.com/manabie-com/rio/internal/database"
	fs "github.com/manabie-com/rio/internal/storage"
)

func ProvideStubStore(ctx context.Context, cfg *config.Config) (rio.StubStore, error) {
	db, err := database.NewStubDBStore(ctx, cfg.DB)
	if err != nil {
		return nil, err
	}

	return cache.NewStubCache(db, db, cfg), nil
}

func ProvideFileStorage(ctx context.Context, cfg *config.Config) (fs.FileStorage, error) {
	return fs.NewFileStorage(cfg.FileStorage), nil
}
