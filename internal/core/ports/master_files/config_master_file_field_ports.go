package ports

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"github.com/billowdev/exclusive-go-hexa/pkg/utils"
)

type IConfigSystemMasterFileFieldRepository interface {
	GetConfigSystemMasterFileField(ctx context.Context, id uint) (*models.ConfigSystemMasterFileField, error)
	GetConfigSystemMasterFileFields(ctx context.Context) (*pagination.Pagination[[]models.ConfigSystemMasterFileField], error)
	CreateConfigSystemMasterFileField(ctx context.Context, payload *models.ConfigSystemMasterFileField) error
	UpdateConfigSystemMasterFileField(ctx context.Context, payload *models.ConfigSystemMasterFileField) error
	DeleteConfigSystemMasterFileField(ctx context.Context, id uint) error
}

type IConfigSystemMasterFileFieldService interface {
	GetConfigSystemMasterFileField(ctx context.Context, id uint) utils.APIResponse
	GetConfigSystemMasterFileFields(ctx context.Context) pagination.Pagination[[]models.ConfigSystemMasterFileField]
	CreateConfigSystemMasterFileField(ctx context.Context, payload *models.ConfigSystemMasterFileField) utils.APIResponse
	UpdateConfigSystemMasterFileField(ctx context.Context, payload *models.ConfigSystemMasterFileField) utils.APIResponse
	DeleteConfigSystemMasterFileField(ctx context.Context, id uint) utils.APIResponse
}
