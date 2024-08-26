package ports

import (
	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/filters"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/pagination"
	"github.com/billowdev/document-system-field-manager/pkg/utils"
)

type ILogMasterFileRepository interface {
	GetLogMasterFile(id uint) (*models.LogMasterFile, error)
	GetLogMasterFiles(p pagination.PaginationParams[filters.LogMasterFileFilter]) (*pagination.Pagination[[]models.LogMasterFile], error)
	CreateLogMasterFile(payload *models.LogMasterFile) error
	UpdateLogMasterFile(payload *models.LogMasterFile) error
	DeleteLogMasterFile(id uint) error
}

type ILogMasterFileService interface {
	GetLogMasterFile(id uint) utils.APIResponse
	GetLogMasterFiles(p pagination.PaginationParams[filters.LogMasterFileFilter]) pagination.Pagination[[]models.LogMasterFile]
	CreateLogMasterFile(payload *models.LogMasterFile) utils.APIResponse
	UpdateLogMasterFile(payload *models.LogMasterFile) utils.APIResponse
	DeleteLogMasterFile(id uint) utils.APIResponse
}
