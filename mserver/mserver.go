package mserver

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/manabie-com/rio/internal/api"
	"github.com/manabie-com/rio/internal/config"
	"github.com/manabie-com/rio/internal/log"
)

func StartInMemoryServer() error {
	gin.SetMode(gin.ReleaseMode)
	api.SetupContext()

	ctx := context.Background()
	app, err := api.NewInMemoryApp(ctx, config.NewConfig())
	if err != nil {
		log.Error(ctx, err)
		return err
	}

	if err := app.Start(ctx); err != nil {
		log.Error(ctx, err)
		return err
	}
	return nil
}
