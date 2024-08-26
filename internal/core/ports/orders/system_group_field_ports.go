package ports

import (
	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/filters"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/pagination"
	"github.com/billowdev/document-system-field-manager/pkg/utils"
)

type ISystemGroupFieldRepository interface {
	GetSystemGroupFields(id uint) (*models.SystemGroupField, error)
	GetSystemGroupFieldss(param pagination.PaginationParams[filters.SystemGroupFieldFilter]) (*pagination.Pagination[[]models.SystemGroupField], error)
	CreateSystemGroupFields(payload *models.SystemGroupField) error
	UpdateSystemGroupFields(payload *models.SystemGroupField) error
	DeleteSystemGroupFields(id uint) error
}

type ISystemGroupFieldService interface {
	GetSystemGroupField(id uint) utils.APIResponse
	GetSystemGroupFields(param pagination.PaginationParams[filters.SystemGroupFieldFilter]) pagination.Pagination[[]models.SystemGroupField]
	CreateSystemGroupField(payload *models.SystemGroupField) utils.APIResponse
	UpdateSystemGroupField(payload *models.SystemGroupField) utils.APIResponse
	DeleteSystemGroupField(id uint) utils.APIResponse
}
