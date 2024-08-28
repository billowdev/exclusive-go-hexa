package main

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/temporal/worker"
	"github.com/billowdev/exclusive-go-hexa/pkg/configs"
	"go.temporal.io/sdk/client"
	temporalLog "go.temporal.io/sdk/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func condb() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v search_path=%v",
		configs.DB_HOST,
		configs.DB_USERNAME,
		configs.DB_PASSWORD,
		configs.DB_NAME,
		configs.DB_PORT,
		configs.DB_SSL_MODE,
		configs.DB_SCHEMA,
	)
	loggerDBLevel := logger.Silent
	if configs.APP_DEBUG_MODE {
		loggerDBLevel = logger.Info
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		// PreferSimpleProtocol: DB_DRY_RUN,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(loggerDBLevel), // or logger.Silent if you don't want logs
		// Logger: logger.Default.LogMode(logger.Info), // or logger.Silent if you don't want logs
	})

	if err != nil {
		return nil, err // instead of panic, return the error
	}
	return db, nil
}
func main() {
	logger := temporalLog.NewStructuredLogger(slog.Default())
	hostPort := client.DefaultHostPort
	if configs.TEMPORAL_CLIENT_URL != "" {
		hostPort = configs.TEMPORAL_CLIENT_URL
	}

	db, err := condb()
	if err != nil {
		log.Fatal("Failed to start Database:", err)
	}

	temporalClient, err := client.Dial(client.Options{
		HostPort: hostPort,
		Logger:   logger,
	})

	worker.RegisterTemporalWorkflow(temporalClient, db)
	if err != nil {
		log.Fatal("Failed to start Temporal worker:", err)
	}

	defer temporalClient.Close()

}
