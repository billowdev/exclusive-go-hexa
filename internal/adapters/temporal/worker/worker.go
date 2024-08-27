package worker

import (
	"log"
	"log/slog"

	repositories "github.com/billowdev/exclusive-go-hexa/internal/adapters/repositories/activity"
	"github.com/billowdev/exclusive-go-hexa/internal/adapters/temporal/activities"
	"github.com/billowdev/exclusive-go-hexa/internal/adapters/temporal/workflows"
	"github.com/billowdev/exclusive-go-hexa/pkg/configs"
	"go.temporal.io/sdk/client"
	temporalLog "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"
	"gorm.io/gorm"
)

func WorkflowClient() client.Client {
	logger := temporalLog.NewStructuredLogger(slog.Default())
	hostPort := func() string {
		if configs.TEMPORAL_CLIENT_URL != "" {
			return configs.TEMPORAL_CLIENT_URL
		}
		return client.DefaultHostPort
	}()

	c, err := client.Dial(client.Options{
		// HostPort: client.DefaultHostPort,
		HostPort: hostPort,
		Logger:   logger,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}

	return c
}
func RegisterTemporalWorkflow(c client.Client, db *gorm.DB) {
	w := worker.New(c, "email_worker", worker.Options{})

	w.RegisterWorkflow(workflows.LoggingWorkflow)

	repo := repositories.NewActivityRepository(db)
	loggingActivity := &activities.LoggingActivity{Repo: repo}

	w.RegisterActivity(loggingActivity)

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln("Unable to start email_worker", err)
	}
	// return nil
}
