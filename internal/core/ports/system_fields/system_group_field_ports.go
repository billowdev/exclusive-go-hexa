package ports

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	domain "github.com/billowdev/exclusive-go-hexa/internal/core/domain/system_fields"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"github.com/billowdev/exclusive-go-hexa/pkg/utils"
)

type ISystemGroupFieldRepository interface {
	GetSystemGroupField(ctx context.Context, id uint) (*models.SystemGroupField, error)
	GetSystemGroupFields(ctx context.Context) (*pagination.Pagination[[]models.SystemGroupField], error)
	CreateSystemGroupField(ctx context.Context, payload *models.SystemGroupField) error
	UpdateSystemGroupField(ctx context.Context, payload *models.SystemGroupField) error
	DeleteSystemGroupField(ctx context.Context, id uint) error
}

type ISystemGroupFieldService interface {
	GetSystemGroupField(ctx context.Context, id uint) utils.APIResponse
	GetSystemGroupFields(ctx context.Context) pagination.Pagination[[]domain.SystemGroupFieldDomain]
	CreateSystemGroupField(ctx context.Context, payload *models.SystemGroupField) utils.APIResponse
	UpdateSystemGroupField(ctx context.Context, id uint, payload *models.SystemGroupField) utils.APIResponse
	DeleteSystemGroupField(ctx context.Context, id uint) utils.APIResponse
}
