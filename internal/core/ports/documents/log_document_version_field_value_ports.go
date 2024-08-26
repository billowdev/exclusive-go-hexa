package ports

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"github.com/billowdev/exclusive-go-hexa/pkg/utils"
)

type ILogDocumentVersionFieldValueRepository interface {
	GetLogDocumentVersionFieldValue(ctx context.Context, id uint) (*models.LogDocumentVersionFieldValue, error)
	GetLogDocumentVersionFieldValues(ctx context.Context) (*pagination.Pagination[[]models.LogDocumentVersionFieldValue], error)
	CreateLogDocumentVersionFieldValue(ctx context.Context, payload *models.LogDocumentVersionFieldValue) error
	UpdateLogDocumentVersionFieldValue(ctx context.Context, payload *models.LogDocumentVersionFieldValue) error
	DeleteLogDocumentVersionFieldValue(ctx context.Context, id uint) error
}

type ILogDocumentVersionFieldValueService interface {
	GetLogDocumentVersionFieldValue(ctx context.Context, id uint) utils.APIResponse
	GetLogDocumentVersionFieldValues(ctx context.Context) pagination.Pagination[[]models.LogDocumentVersionFieldValue]
	CreateLogDocumentVersionFieldValue(ctx context.Context, payload *models.LogDocumentVersionFieldValue) utils.APIResponse
	UpdateLogDocumentVersionFieldValue(ctx context.Context, payload *models.LogDocumentVersionFieldValue) utils.APIResponse
	DeleteLogDocumentVersionFieldValue(ctx context.Context, id uint) utils.APIResponse
}
