package workflows

import (
	"time"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/temporal/activities"
	"github.com/billowdev/exclusive-go-hexa/internal/adapters/temporal/dto"
	"go.temporal.io/sdk/workflow"
)

func RegistrationWorkflow(ctx workflow.Context, data dto.RegistrationData) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second, // Adjust based on your needs
		// RetryPolicy: &internal.RetryPolicy{

		// },
	}

	ctx = workflow.WithActivityOptions(ctx, ao)

	// Step 1: Register Data
	// var registerResult *pagination.Pagination[[]models.SystemField]

	// err := workflow.ExecuteActivity(ctx, activities.RegisterDataActivity, data).Get(ctx, &registerResult)
	err := workflow.ExecuteActivity(ctx, activities.RegisterDataActivity, data).Get(ctx, nil)
	if err != nil {
		return err
	}

	// Step 2: Check Data in System
	var status string
	err = workflow.ExecuteActivity(ctx, activities.CheckDataActivity, data).Get(ctx, &status)
	if err != nil {
		return err
	}

	if status == "pending" {
		// Step 3: Verify Identities
		err = workflow.ExecuteActivity(ctx, activities.VerifyIdentitiesActivity, data).Get(ctx, nil)
		if err != nil {
			return err
		}

		// Step 4: Update Registration Status
		err = workflow.ExecuteActivity(ctx, activities.UpdateStatusActivity, "success").Get(ctx, nil)
		if err != nil {
			return err
		}
	} else {
		return workflow.NewContinueAsNewError(ctx, "Data not in pending status")
	}

	return nil
}
