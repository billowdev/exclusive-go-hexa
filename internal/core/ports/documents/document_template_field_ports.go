package ports

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"github.com/billowdev/exclusive-go-hexa/pkg/utils"
)

type IDocumentTemplateFieldRepository interface {
	GetDocumentTemplateField(ctx context.Context, id uint) (*models.DocumentTemplateField, error)
	GetDocumentTemplateFields(ctx context.Context) (*pagination.Pagination[[]models.DocumentTemplateField], error)
	CreateDocumentTemplateField(ctx context.Context, payload *models.DocumentTemplateField) error
	UpdateDocumentTemplateField(ctx context.Context, payload *models.DocumentTemplateField) error
	DeleteDocumentTemplateField(ctx context.Context, id uint) error
}

type IDocumentTemplateFieldService interface {
	GetDocumentTemplateField(ctx context.Context, id uint) utils.APIResponse
	GetDocumentTemplateFields(ctx context.Context) pagination.Pagination[[]models.DocumentTemplateField]
	CreateDocumentTemplateField(ctx context.Context, payload *models.DocumentTemplateField) utils.APIResponse
	UpdateDocumentTemplateField(ctx context.Context, payload *models.DocumentTemplateField) utils.APIResponse
	DeleteDocumentTemplateField(ctx context.Context, id uint) utils.APIResponse
}
