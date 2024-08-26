package ports

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/filters"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"github.com/billowdev/exclusive-go-hexa/pkg/utils"
)

type ISystemGroupFieldRepository interface {
	GetSystemGroupFields(ctx context.Context, id uint) (*models.SystemGroupField, error)
	GetSystemGroupFieldss(ctx context.Context, p pagination.PaginationParams[filters.SystemGroupFieldFilter]) (*pagination.Pagination[[]models.SystemGroupField], error)
	CreateSystemGroupFields(ctx context.Context, payload *models.SystemGroupField) error
	UpdateSystemGroupFields(ctx context.Context, payload *models.SystemGroupField) error
	DeleteSystemGroupFields(ctx context.Context, id uint) error
}

type ISystemGroupFieldService interface {
	GetSystemGroupField(ctx context.Context, id uint) utils.APIResponse
	GetSystemGroupFields(ctx context.Context, p pagination.PaginationParams[filters.SystemGroupFieldFilter]) pagination.Pagination[[]models.SystemGroupField]
	CreateSystemGroupField(ctx context.Context, payload *models.SystemGroupField) utils.APIResponse
	UpdateSystemGroupField(ctx context.Context, payload *models.SystemGroupField) utils.APIResponse
	DeleteSystemGroupField(ctx context.Context, id uint) utils.APIResponse
}
