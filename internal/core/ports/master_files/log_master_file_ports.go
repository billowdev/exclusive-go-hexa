package ports

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"github.com/billowdev/exclusive-go-hexa/pkg/utils"
)

type ILogMasterFileRepository interface {
	GetLogMasterFile(ctx context.Context, id uint) (*models.LogMasterFile, error)
	GetLogMasterFiles(ctx context.Context) (*pagination.Pagination[[]models.LogMasterFile], error)
	CreateLogMasterFile(ctx context.Context, payload *models.LogMasterFile) error
	UpdateLogMasterFile(ctx context.Context, payload *models.LogMasterFile) error
	DeleteLogMasterFile(ctx context.Context, id uint) error
}

type ILogMasterFileService interface {
	GetLogMasterFile(ctx context.Context, id uint) utils.APIResponse
	GetLogMasterFiles(ctx context.Context) pagination.Pagination[[]models.LogMasterFile]
	CreateLogMasterFile(ctx context.Context, payload *models.LogMasterFile) utils.APIResponse
	UpdateLogMasterFile(ctx context.Context, payload *models.LogMasterFile) utils.APIResponse
	DeleteLogMasterFile(ctx context.Context, id uint) utils.APIResponse
}
