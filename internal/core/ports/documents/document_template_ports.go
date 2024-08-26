package ports

import (
	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/filters"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/pagination"
	"github.com/billowdev/document-system-field-manager/pkg/utils"
)

type IDocumentTemplateRepository interface {
	GetDocumentTemplate(id uint) (*models.DocumentTemplate, error)
	GetDocumentTemplates(param pagination.PaginationParams[filters.DocumentTemplateFilter]) (*pagination.Pagination[[]models.DocumentTemplate], error)
	CreateDocumentTemplate(payload *models.DocumentTemplate) error
	UpdateDocumentTemplate(payload *models.DocumentTemplate) error
	DeleteDocumentTemplate(id uint) error
}

type IDocumentTemplateService interface {
	GetDocumentTemplate(id uint) utils.APIResponse
	GetDocumentTemplates(param pagination.PaginationParams[filters.DocumentTemplateFilter]) pagination.Pagination[[]models.DocumentTemplate]
	CreateDocumentTemplate(payload *models.DocumentTemplate) utils.APIResponse
	UpdateDocumentTemplate(payload *models.DocumentTemplate) utils.APIResponse
	DeleteDocumentTemplate(id uint) utils.APIResponse
}
