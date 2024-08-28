package activities

import (
	"context"
	"time"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	ports "github.com/billowdev/exclusive-go-hexa/internal/core/ports/activities"
)

type LoggingActivity struct {
	Repo ports.IActivityRepository
}

type LogActivityInput struct {
	WorkflowID string                 `json:"workflow_id"`
	Name       string                 `json:"name"`
	Status     string                 `json:"status"`
	Input      map[string]interface{} `json:"input"`
	Output     map[string]interface{} `json:"output"`
	Error      string                 `json:"error"`
	StartTime  time.Time              `json:"start_time"`
	EndTime    time.Time              `json:"end_time"`
}

func (a *LoggingActivity) LogActivity(ctx context.Context, input LogActivityInput) error {
	activity := models.Activity{
		WorkflowID: input.WorkflowID,
		Name:       input.Name,
		Status:     input.Status,
		Input:      input.Input,
		Output:     input.Output,
		Error:      input.Error,
		StartTime:  input.StartTime,
		EndTime:    input.EndTime,
	}
	return a.Repo.CreateActivity(ctx, &activity)
}
