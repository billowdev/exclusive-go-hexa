package ports

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"github.com/billowdev/exclusive-go-hexa/pkg/utils"
)

type IDocumentVersionFieldValueRepository interface {
	GetDocumentVersionFieldValue(ctx context.Context, id uint) (*models.DocumentVersionFieldValue, error)
	GetDocumentVersionFieldValues(ctx context.Context) (*pagination.Pagination[[]models.DocumentVersionFieldValue], error)
	CreateDocumentVersionFieldValue(ctx context.Context, payload *models.DocumentVersionFieldValue) error
	UpdateDocumentVersionFieldValue(ctx context.Context, payload *models.DocumentVersionFieldValue) error
	DeleteDocumentVersionFieldValue(ctx context.Context, id uint) error
}

type IDocumentVersionFieldValueService interface {
	GetDocumentVersionFieldValue(ctx context.Context, id uint) utils.APIResponse
	GetDocumentVersionFieldValues(ctx context.Context) pagination.Pagination[[]models.DocumentVersionFieldValue]
	CreateDocumentVersionFieldValue(ctx context.Context, payload *models.DocumentVersionFieldValue) utils.APIResponse
	UpdateDocumentVersionFieldValue(ctx context.Context, payload *models.DocumentVersionFieldValue) utils.APIResponse
	DeleteDocumentVersionFieldValue(ctx context.Context, id uint) utils.APIResponse
}
