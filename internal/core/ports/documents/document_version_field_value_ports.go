package ports

import (
	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/filters"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/pagination"
	"github.com/billowdev/document-system-field-manager/pkg/utils"
)

type IDocumentVersionFieldValueRepository interface {
	GetDocumentVersionFieldValue(id uint) (*models.DocumentVersionFieldValue, error)
	GetDocumentVersionFieldValues(param pagination.PaginationParams[filters.DocumentVersionFieldValueFilter]) (*pagination.Pagination[[]models.DocumentVersionFieldValue], error)
	CreateDocumentVersionFieldValue(payload *models.DocumentVersionFieldValue) error
	UpdateDocumentVersionFieldValue(payload *models.DocumentVersionFieldValue) error
	DeleteDocumentVersionFieldValue(id uint) error
}

type IDocumentVersionFieldValueService interface {
	GetDocumentVersionFieldValue(id uint) utils.APIResponse
	GetDocumentVersionFieldValues(param pagination.PaginationParams[filters.DocumentVersionFieldValueFilter]) pagination.Pagination[[]models.DocumentVersionFieldValue]
	CreateDocumentVersionFieldValue(payload *models.DocumentVersionFieldValue) utils.APIResponse
	UpdateDocumentVersionFieldValue(payload *models.DocumentVersionFieldValue) utils.APIResponse
	DeleteDocumentVersionFieldValue(id uint) utils.APIResponse
}
