package ports

import (
	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/filters"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/pagination"
	"github.com/billowdev/document-system-field-manager/pkg/utils"
)

type IDocumentRepository interface {
	GetDocument(id uint) (*models.Document, error)
	GetDocuments(param pagination.PaginationParams[filters.DocumentFilter]) (*pagination.Pagination[[]models.Document], error)
	CreateDocument(payload *models.Document) error
	UpdateDocument(payload *models.Document) error
	DeleteDocument(id uint) error
}

type IDocumentService interface {
	GetDocument(id uint) utils.APIResponse
	GetDocuments(param pagination.PaginationParams[filters.DocumentFilter]) pagination.Pagination[[]models.Document]
	CreateDocument(payload *models.Document) utils.APIResponse
	UpdateDocument(payload *models.Document) utils.APIResponse
	DeleteDocument(id uint) utils.APIResponse
}
