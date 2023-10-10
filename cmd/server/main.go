package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/manabie-com/rio/internal/api"
	"github.com/manabie-com/rio/internal/config"
	"github.com/manabie-com/rio/internal/database"
	"github.com/manabie-com/rio/internal/log"
)

// @title       HTTP Mock
// @description A mock framework for unit test http in golang, also support integration test
// @version     1.0
// @BasePath    /api/v1
func main() {
	gin.SetMode(gin.ReleaseMode)
	api.SetupContext()

	ctx := context.Background()
	app, err := api.NewApp(ctx, config.NewConfig())
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

func StartWithConfig(cfg *config.Config) {
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
