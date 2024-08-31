package ports

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	domain "github.com/billowdev/exclusive-go-hexa/internal/core/domain/system_fields"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"github.com/billowdev/exclusive-go-hexa/pkg/utils"
)

type ISystemFieldRepository interface {
	GetSystemField(ctx context.Context, id uint) (*models.SystemField, error)
	GetSystemFields(ctx context.Context) (*pagination.Pagination[[]models.SystemField], error)
	CreateSystemField(ctx context.Context, payload *models.SystemField) error
	UpdateSystemField(ctx context.Context, payload *models.SystemField) error
	DeleteSystemField(ctx context.Context, id uint) error
}

type ISystemFieldService interface {
	GetSystemField(ctx context.Context, id uint) utils.APIResponse
	GetSystemFields(ctx context.Context) pagination.Pagination[[]domain.SystemFieldDomain]
	CreateSystemField(ctx context.Context, payload *models.SystemField) utils.APIResponse
	UpdateSystemField(ctx context.Context, id uint, payload *models.SystemField) utils.APIResponse
	DeleteSystemField(ctx context.Context, id uint) utils.APIResponse
}
