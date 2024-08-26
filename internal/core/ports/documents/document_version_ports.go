package ports

import (
	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/filters"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/pagination"
	"github.com/billowdev/document-system-field-manager/pkg/utils"
)

type IDocumentVersionRepository interface {
	GetDocumentVersion(id uint) (*models.DocumentVersion, error)
	GetDocumentVersions(param pagination.PaginationParams[filters.DocumentVersionFilter]) (*pagination.Pagination[[]models.DocumentVersion], error)
	CreateDocumentVersion(payload *models.DocumentVersion) error
	UpdateDocumentVersion(payload *models.DocumentVersion) error
	DeleteDocumentVersion(id uint) error
}

type IDocumentVersionService interface {
	GetDocumentVersion(id uint) utils.APIResponse
	GetDocumentVersions(param pagination.PaginationParams[filters.DocumentVersionFilter]) pagination.Pagination[[]models.DocumentVersion]
	CreateDocumentVersion(payload *models.DocumentVersion) utils.APIResponse
	UpdateDocumentVersion(payload *models.DocumentVersion) utils.APIResponse
	DeleteDocumentVersion(id uint) utils.APIResponse
}
