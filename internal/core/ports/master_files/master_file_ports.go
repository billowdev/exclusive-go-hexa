package ports

import (
	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/filters"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/pagination"
	"github.com/billowdev/document-system-field-manager/pkg/utils"
)

type IMasterFileRepository interface {
	GetMasterFile(id uint) (*models.MasterFile, error)
	GetMasterFiles(p pagination.PaginationParams[filters.MasterFileFilter]) (*pagination.Pagination[[]models.MasterFile], error)
	CreateMasterFile(payload *models.MasterFile) error
	UpdateMasterFile(payload *models.MasterFile) error
	DeleteMasterFile(id uint) error
}

type IMasterFileService interface {
	GetMasterFile(id uint) utils.APIResponse
	GetMasterFiles(p pagination.PaginationParams[filters.MasterFileFilter]) pagination.Pagination[[]models.MasterFile]
	CreateMasterFile(payload *models.MasterFile) utils.APIResponse
	UpdateMasterFile(payload *models.MasterFile) utils.APIResponse
	DeleteMasterFile(id uint) utils.APIResponse
}
