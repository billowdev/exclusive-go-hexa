package ports

import (
	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/filters"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/pagination"
	"github.com/billowdev/document-system-field-manager/pkg/utils"
	"gorm.io/gorm"
)

type IDocumentVersionFieldValueRepository interface {
	BeginTransaction() *gorm.DB
	GetDocumentVersionFieldValue(id uint) (*models.DocumentVersionFieldValue, error)
	GetDocumentVersionFieldValues(p pagination.PaginationParams[filters.DocumentVersionFieldValueFilter]) (*pagination.Pagination[[]models.DocumentVersionFieldValue], error)
	CreateDocumentVersionFieldValue(payload *models.DocumentVersionFieldValue) error
	UpdateDocumentVersionFieldValue(payload *models.DocumentVersionFieldValue) error
	DeleteDocumentVersionFieldValue(id uint) error
}

type IDocumentVersionFieldValueService interface {
	GetDocumentVersionFieldValue(id uint) utils.APIResponse
	GetDocumentVersionFieldValues(p pagination.PaginationParams[filters.DocumentVersionFieldValueFilter]) pagination.Pagination[[]models.DocumentVersionFieldValue]
	CreateDocumentVersionFieldValue(payload *models.DocumentVersionFieldValue) utils.APIResponse
	UpdateDocumentVersionFieldValue(payload *models.DocumentVersionFieldValue) utils.APIResponse
	DeleteDocumentVersionFieldValue(id uint) utils.APIResponse
}
