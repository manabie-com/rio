package mserver

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/manabie-com/rio/internal/api"
	"github.com/manabie-com/rio/internal/config"
	"github.com/manabie-com/rio/internal/database"
	"github.com/manabie-com/rio/internal/log"
	fs "github.com/manabie-com/rio/internal/storage"
)

func StartServerWithConfig(cfg *config.Config) error {
	gin.SetMode(gin.ReleaseMode)
	api.SetupContext()

	ctx := context.Background()
	app, err := api.NewApp(ctx, cfg)
	if err != nil {
		log.Error(ctx, err)
		return err
	}

	dbConfig := cfg.DB
	if err := database.Migrate(ctx, dbConfig, "schema/migration"); err != nil {
		log.Error(ctx, err)
		return err
	}

	if err := app.Start(ctx); err != nil {
		log.Error(ctx, err)
		return err
	}
	return nil
}

// NewConfig
// serverAddress: 0.0.0.0:8896
func NewConfig(serverAddr, dbHost, dbName, dbSchema, dbUser, dbPassword, dbPort string) *config.Config {
	return &config.Config{
		ServerAddress:   serverAddr,
		DB:              NewDBConfig(dbHost, dbName, dbSchema, dbUser, dbPassword, dbPort, 20, 2, 4),
		FileStorageType: "local",
		FileStorage: fs.LocalStorageConfig{
			UseTempDir:  true,
			StoragePath: "uploaded_files",
		},
	}
}

func NewDBConfig(host, dbName, dbSchema, user, password, port string, connectionLifetimeSeconds, maxIdleConnections, maxOpenConnections int) *config.PostgresConfig {
	return &config.PostgresConfig{
		Host:                      host,
		DBName:                    dbName,
		DBSchema:                  dbSchema,
		User:                      user,
		Password:                  password,
		Port:                      port,
		ConnectionLifetimeSeconds: connectionLifetimeSeconds,
		MaxIdleConnections:        maxIdleConnections,
		MaxOpenConnections:        maxOpenConnections,
	}
}
