package workflows

import (
	"time"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/temporal/activities"
	"github.com/billowdev/exclusive-go-hexa/internal/adapters/temporal/dto"
	"go.temporal.io/sdk/workflow"
)

// func RequestWorkflow(ctx workflow.Context, requestData string) (string, error) {
// 	logger := workflow.GetLogger(ctx)
// 	logger.Info("RequestWorkflow started", "RequestData", requestData)

// 	// Set up an activity options
// 	ao := workflow.ActivityOptions{
// 		StartToCloseTimeout: time.Minute,
// 	}
// 	ctx = workflow.WithActivityOptions(ctx, ao)

// 	// Execute the activity
// 	var result string
// 	err := workflow.ExecuteActivity(ctx, activities.ProcessRequestActivity, requestData).Get(ctx, &result)
// 	if err != nil {
// 		logger.Error("Activity failed", "Error", err)
// 		return "", err
// 	}

// 	logger.Info("RequestWorkflow completed", "Result", result)
// 	return result, nil
// }

func TicketPurchaseWorkflow(ctx workflow.Context, request dto.TicketPurchaseRequest) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("TicketPurchaseWorkflow started", "Request", request)

	// Activity options
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// Step 1: Validate the request
	err := workflow.ExecuteActivity(ctx, activities.ValidateRequestActivity, request).Get(ctx, nil)
	if err != nil {
		logger.Error("Failed to validate request", "Error", err)
		return "", err
	}

	// Step 2: Process the payment
	err = workflow.ExecuteActivity(ctx, activities.ProcessPaymentActivity, request).Get(ctx, nil)
	if err != nil {
		logger.Error("Failed to process payment", "Error", err)
		return "", err
	}

	// Step 3: Reserve the ticket
	var reservationID string
	err = workflow.ExecuteActivity(ctx, activities.ReserveTicketActivity, request).Get(ctx, &reservationID)
	if err != nil {
		logger.Error("Failed to reserve ticket", "Error", err)
		return "", err
	}

	// Step 4: Send confirmation
	err = workflow.ExecuteActivity(ctx, activities.SendConfirmationActivity, request, reservationID).Get(ctx, nil)
	if err != nil {
		logger.Error("Failed to send confirmation", "Error", err)
		return "", err
	}

	logger.Info("TicketPurchaseWorkflow completed", "ReservationID", reservationID)
	return reservationID, nil
}
