package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/manabie-com/rio/internal/config"
	"github.com/manabie-com/rio/internal/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect setups connections to Postgres database
func Connect(ctx context.Context, config *config.PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"user='%s' password='%s' host='%s' port='%s' dbname='%s' sslmode='disable' application_name='%s'",
		config.User, config.Password, config.Host, config.Port, config.DBName, config.User,
	)
	cfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}

	db, err := gorm.Open(postgres.Open(dsn), cfg)
	if err != nil {
		log.Error(ctx, "cannot connect to database", config.DBName)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error(ctx, "cannot obtain sql database object", config.DBName)
		return nil, err
	}

	sqlDB.SetConnMaxLifetime(time.Duration(20) * time.Second)
	sqlDB.SetMaxIdleConns(config.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(config.MaxOpenConnections)

	/* Disable tracing causes of the library conflict
	if config.EnableTracing {
		if err = db.WithContext(ctx).Use(otelgorm.NewPlugin(otelgorm.WithDBName(config.Schema))); err != nil {
			log.Error(ctx, "cannot enable tracing for gorm", config.Schema)
			return nil, err
		}
	}
	*/

	log.Info(ctx, "connected to database", config.DBName)
	return db, nil
}

// Disconnect closes the connections to the MySQL database
func Disconnect(ctx context.Context, db *gorm.DB) error {
	sqlDB, err := db.WithContext(ctx).DB()
	if err != nil {
		log.Error(ctx, err)
		return err
	}

	if err := sqlDB.Close(); err != nil {
		log.Error(ctx, "cannot close", err)
		return err
	}

	return nil
}

// ExecuteFileScript runs a specific migration file on a MySQL database using specific path
func ExecuteFileScript(ctx context.Context, db *gorm.DB, filePath string) error {
	migrationSQL, err := os.ReadFile(filePath)
	if err != nil {
		log.Error(ctx, "cannot read sql script", err, filePath)
		return err
	}

	if err := db.WithContext(ctx).Exec(string(migrationSQL)).Error; err != nil {
		log.Error(ctx, err)
		return err
	}

	log.Info(ctx, "executed sql script", filePath)
	return nil
}
