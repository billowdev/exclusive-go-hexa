package main

import (
	"log"
	"log/slog"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/temporal/worker"
	"github.com/billowdev/exclusive-go-hexa/pkg/configs"
	"go.temporal.io/sdk/client"
	temporalLog "go.temporal.io/sdk/log"
	"gorm.io/gorm"
)

func main() {
	logger := temporalLog.NewStructuredLogger(slog.Default())
	hostPort := client.DefaultHostPort
	if configs.TEMPORAL_CLIENT_URL != "" {
		hostPort = configs.TEMPORAL_CLIENT_URL
	}

	temporalClient, err := client.Dial(client.Options{
		HostPort: hostPort,
		Logger:   logger,
	})

	worker.RegisterTemporalWorkflow(temporalClient, &gorm.DB{})
	if err != nil {
		log.Fatal("Failed to start Temporal worker:", err)
	}

	defer temporalClient.Close()

}
