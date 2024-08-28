package workflows

import (
	"time"

	repositories "github.com/billowdev/exclusive-go-hexa/internal/adapters/repositories/activity"
	"github.com/billowdev/exclusive-go-hexa/internal/adapters/temporal/activities"
	"go.temporal.io/sdk/workflow"
	"gorm.io/gorm"
)

type LoggingWorkflowInput struct {
	WorkflowID string
	Name       string
	Status     string
	Input      map[string]interface{}
	Output     map[string]interface{}
	Error      string
	StartTime  time.Time
	EndTime    time.Time
}

func LoggingWorkflow(ctx workflow.Context, db *gorm.DB, input LoggingWorkflowInput) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	repo := repositories.NewActivityRepository(db)
	loggingActivity := &activities.LoggingActivity{Repo: repo}

	err := workflow.ExecuteActivity(ctx, loggingActivity, input).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
