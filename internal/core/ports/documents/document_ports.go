package ports

import (
	"context"

	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/pagination"
	"github.com/billowdev/document-system-field-manager/pkg/utils"
)

type IDocumentRepository interface {
	GetDocument(ctx context.Context, id uint) (*models.Document, error)
	GetDocuments(ctx context.Context) (*pagination.Pagination[[]models.Document], error)
	CreateDocument(ctx context.Context, payload *models.Document) error
	UpdateDocument(ctx context.Context, payload *models.Document) error
	DeleteDocument(ctx context.Context, id uint) error
}

type IDocumentService interface {
	GetDocument(ctx context.Context, id uint) utils.APIResponse
	GetDocuments(ctx context.Context) pagination.Pagination[[]models.Document]
	CreateDocument(ctx context.Context, payload *models.Document) utils.APIResponse
	UpdateDocument(ctx context.Context, payload *models.Document) utils.APIResponse
	DeleteDocument(ctx context.Context, id uint) utils.APIResponse
}
