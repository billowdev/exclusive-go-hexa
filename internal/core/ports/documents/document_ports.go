package ports

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"github.com/billowdev/exclusive-go-hexa/pkg/utils"
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
