package ports

import (
	"context"

	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/filters"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/pagination"
	"github.com/billowdev/document-system-field-manager/pkg/utils"
)

type ILogMasterFileRepository interface {
	GetLogMasterFile(ctx context.Context, id uint) (*models.LogMasterFile, error)
	GetLogMasterFiles(ctx context.Context, p pagination.PaginationParams[filters.LogMasterFileFilter]) (*pagination.Pagination[[]models.LogMasterFile], error)
	CreateLogMasterFile(ctx context.Context, payload *models.LogMasterFile) error
	UpdateLogMasterFile(ctx context.Context, payload *models.LogMasterFile) error
	DeleteLogMasterFile(ctx context.Context, id uint) error
}

type ILogMasterFileService interface {
	GetLogMasterFile(ctx context.Context, id uint) utils.APIResponse
	GetLogMasterFiles(ctx context.Context, p pagination.PaginationParams[filters.LogMasterFileFilter]) pagination.Pagination[[]models.LogMasterFile]
	CreateLogMasterFile(ctx context.Context, payload *models.LogMasterFile) utils.APIResponse
	UpdateLogMasterFile(ctx context.Context, payload *models.LogMasterFile) utils.APIResponse
	DeleteLogMasterFile(ctx context.Context, id uint) utils.APIResponse
}
