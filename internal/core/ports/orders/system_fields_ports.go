package ports

import (
	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/filters"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/pagination"
	"github.com/billowdev/document-system-field-manager/pkg/utils"
)

type ISystemFieldRepository interface {
	GetSystemField(id uint) (*models.SystemField, error)
	GetSystemFields(param pagination.PaginationParams[filters.SystemFieldFilter]) (*pagination.Pagination[[]models.SystemField], error)
	CreateSystemField(payload *models.SystemField) error
	UpdateSystemField(payload *models.SystemField) error
	DeleteSystemField(id uint) error
}

type ISystemFieldService interface {
	GetSystemField(id uint) utils.APIResponse
	GetSystemFields(param pagination.PaginationParams[filters.SystemFieldFilter]) pagination.Pagination[[]models.SystemField]
	CreateSystemField(payload *models.SystemField) utils.APIResponse
	UpdateSystemField(payload *models.SystemField) utils.APIResponse
	DeleteSystemField(id uint) utils.APIResponse
}
