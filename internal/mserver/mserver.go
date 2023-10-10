package mserver

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/manabie-com/rio/internal/api"
	"github.com/manabie-com/rio/internal/config"
	"github.com/manabie-com/rio/internal/database"
	"github.com/manabie-com/rio/internal/log"
)

func StartServerWithConfig(cfg *config.Config) {
	gin.SetMode(gin.ReleaseMode)
	api.SetupContext()

	ctx := context.Background()
	app, err := api.NewApp(ctx, cfg)
	if err != nil {
		log.Error(ctx, err)
		panic(err)
	}

	dbConfig := config.NewDBConfig()
	if err := database.Migrate(ctx, dbConfig, "schema/migration"); err != nil {
		panic(err)
	}

	if err := app.Start(ctx); err != nil {
		log.Error(ctx, err)
		panic(err)
	}
}
