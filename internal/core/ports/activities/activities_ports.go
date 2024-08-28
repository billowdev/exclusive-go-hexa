package ports

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
)

type IActivityRepository interface {
	CreateActivity(ctx context.Context, activity *models.Activity) error
}
