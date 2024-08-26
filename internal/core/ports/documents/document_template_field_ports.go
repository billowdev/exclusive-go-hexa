package ports

import (
	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/filters"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/pagination"
	"github.com/billowdev/document-system-field-manager/pkg/utils"
)

type IDocumentTemplateFieldRepository interface {
	GetDocumentTemplateField(id uint) (*models.DocumentTemplateField, error)
	GetDocumentTemplateFields(param pagination.PaginationParams[filters.DocumentTemplateFieldFilter]) (*pagination.Pagination[[]models.DocumentTemplateField], error)
	CreateDocumentTemplateField(payload *models.DocumentTemplateField) error
	UpdateDocumentTemplateField(payload *models.DocumentTemplateField) error
	DeleteDocumentTemplateField(id uint) error
}

type IDocumentTemplateFieldService interface {
	GetDocumentTemplateField(id uint) utils.APIResponse
	GetDocumentTemplateFields(param pagination.PaginationParams[filters.DocumentTemplateFieldFilter]) pagination.Pagination[[]models.DocumentTemplateField]
	CreateDocumentTemplateField(payload *models.DocumentTemplateField) utils.APIResponse
	UpdateDocumentTemplateField(payload *models.DocumentTemplateField) utils.APIResponse
	DeleteDocumentTemplateField(id uint) utils.APIResponse
}
