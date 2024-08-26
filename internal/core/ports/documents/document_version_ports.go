package ports

import (
	"context"

	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/pagination"
	"github.com/billowdev/document-system-field-manager/pkg/utils"
)

type IDocumentVersionRepository interface {
	GetDocumentVersion(ctx context.Context, id uint) (*models.DocumentVersion, error)
	GetDocumentVersions(ctx context.Context) (*pagination.Pagination[[]models.DocumentVersion], error)
	CreateDocumentVersion(ctx context.Context, payload *models.DocumentVersion) error
	UpdateDocumentVersion(ctx context.Context, payload *models.DocumentVersion) error
	DeleteDocumentVersion(ctx context.Context, id uint) error
}

type IDocumentVersionService interface {
	GetDocumentVersion(ctx context.Context, id uint) utils.APIResponse
	GetDocumentVersions(ctx context.Context) pagination.Pagination[[]models.DocumentVersion]
	CreateDocumentVersion(ctx context.Context, payload *models.DocumentVersion) utils.APIResponse
	UpdateDocumentVersion(ctx context.Context, payload *models.DocumentVersion) utils.APIResponse
	DeleteDocumentVersion(ctx context.Context, id uint) utils.APIResponse
}
