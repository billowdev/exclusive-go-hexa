package ports

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"github.com/billowdev/exclusive-go-hexa/pkg/utils"
)

type IActivityRepository interface {
	GetActivity(ctx context.Context, id uint) (*models.Activity, error)
	GetActivities(ctx context.Context) (*pagination.Pagination[[]models.Activity], error)
	CreateActivity(ctx context.Context, payload *models.Activity) error
	UpdateActivity(ctx context.Context, payload *models.Activity) error
	DeleteActivity(ctx context.Context, id uint) error
}

type IActivityService interface {
	GetActivity(ctx context.Context, id uint) utils.APIResponse
	GetActivities(ctx context.Context) pagination.Pagination[[]models.Activity]
	CreateActivity(ctx context.Context, payload *models.Activity) utils.APIResponse
	UpdateActivity(ctx context.Context, payload *models.Activity) utils.APIResponse
	DeleteActivity(ctx context.Context, id uint) utils.APIResponse
}
