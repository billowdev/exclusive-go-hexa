package ports

import (
	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/filters"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/pagination"
	"github.com/billowdev/document-system-field-manager/pkg/utils"
)

type ILogDocumentVersionFieldValueRepository interface {
	GetLogDocumentVersionFieldValue(id uint) (*models.LogDocumentVersionFieldValue, error)
	GetLogDocumentVersionFieldValues(param pagination.PaginationParams[filters.LogDocumentVersionFieldValueFilter]) (*pagination.Pagination[[]models.LogDocumentVersionFieldValue], error)
	CreateLogDocumentVersionFieldValue(payload *models.LogDocumentVersionFieldValue) error
	UpdateLogDocumentVersionFieldValue(payload *models.LogDocumentVersionFieldValue) error
	DeleteLogDocumentVersionFieldValue(id uint) error
}

type ILogDocumentVersionFieldValueService interface {
	GetLogDocumentVersionFieldValue(id uint) utils.APIResponse
	GetLogDocumentVersionFieldValues(param pagination.PaginationParams[filters.LogDocumentVersionFieldValueFilter]) pagination.Pagination[[]models.LogDocumentVersionFieldValue]
	CreateLogDocumentVersionFieldValue(payload *models.LogDocumentVersionFieldValue) utils.APIResponse
	UpdateLogDocumentVersionFieldValue(payload *models.LogDocumentVersionFieldValue) utils.APIResponse
	DeleteLogDocumentVersionFieldValue(id uint) utils.APIResponse
}
