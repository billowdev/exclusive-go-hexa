package ports

import (
	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/filters"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/pagination"
	"github.com/billowdev/document-system-field-manager/pkg/utils"
)

type IConfigSystemMasterFileFieldRepository interface {
	GetConfigSystemMasterFileField(id uint) (*models.ConfigSystemMasterFileField, error)
	GetConfigSystemMasterFileFields(param pagination.PaginationParams[filters.ConfigSystemMasterFileFieldFilter]) (*pagination.Pagination[[]models.ConfigSystemMasterFileField], error)
	CreateConfigSystemMasterFileField(payload *models.ConfigSystemMasterFileField) error
	UpdateConfigSystemMasterFileField(payload *models.ConfigSystemMasterFileField) error
	DeleteConfigSystemMasterFileField(id uint) error
}

type IConfigSystemMasterFileFieldService interface {
	GetConfigSystemMasterFileField(id uint) utils.APIResponse
	GetConfigSystemMasterFileFields(param pagination.PaginationParams[filters.ConfigSystemMasterFileFieldFilter]) pagination.Pagination[[]models.ConfigSystemMasterFileField]
	CreateConfigSystemMasterFileField(payload *models.ConfigSystemMasterFileField) utils.APIResponse
	UpdateConfigSystemMasterFileField(payload *models.ConfigSystemMasterFileField) utils.APIResponse
	DeleteConfigSystemMasterFileField(id uint) utils.APIResponse
}
