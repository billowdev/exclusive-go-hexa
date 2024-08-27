package worker

import (
	"log"
	"log/slog"

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
	workerOptions := worker.Options{
		// MaxConcurrentActivityExecutionSize: 1, // Limits the number of concurrent activities
		// MaxConcurrentWorkflowTaskPollers:   1, // Ensures only one workflow task poller
	}
	ticketWorker := worker.New(c, "TICKET_TASK_QUEUE", workerOptions)
	ticketWorker.RegisterWorkflow(workflows.TicketPurchaseWorkflow)
	ticketWorker.RegisterActivity(activities.ValidateRequestActivity)
	ticketWorker.RegisterActivity(activities.ProcessPaymentActivity)
	ticketWorker.RegisterActivity(activities.ReserveTicketActivity)
	ticketWorker.RegisterActivity(activities.SendConfirmationActivity)

	registerWorker := worker.New(c, "REGISTER_TASK_QUEUE", workerOptions)
	registerWorker.RegisterWorkflow(workflows.RegistrationWorkflow)
	registerWorker.RegisterActivity(activities.RegisterDataActivity)
	registerWorker.RegisterActivity(activities.CheckDataActivity)
	registerWorker.RegisterActivity(activities.VerifyIdentitiesActivity)
	registerWorker.RegisterActivity(activities.UpdateStatusActivity)

	// go func() {
	// 	if err := ticketWorker.Run(worker.InterruptCh()); err != nil {
	// 		log.Fatalf("Failed to start ticketWorker: %v", err)
	// 	}
	// }()
	// go func() {
	// 	if err := registerWorker.Run(worker.InterruptCh()); err != nil {
	// 		log.Fatalf("Failed to start registerWorker: %v", err)
	// 	}
	// }()
	// select {}
	err := ticketWorker.Start()
	if err != nil {
		log.Fatalf("Failed to start worker: %v", err)
	}
	err = registerWorker.Start()
	if err != nil {
		log.Fatalf("Failed to start worker: %v", err)
	}
	select {}
	// if err := registerWorker.Run(worker.InterruptCh()); err != nil {
	// 	log.Fatalf("Failed to start registerWorker: %v", err)
	// }

}
