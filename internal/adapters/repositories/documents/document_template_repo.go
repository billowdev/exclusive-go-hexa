package repositories

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database"
	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	ports "github.com/billowdev/exclusive-go-hexa/internal/core/ports/documents"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/filters"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"gorm.io/gorm"
)

type DocumentTemplateImpl struct {
	db *gorm.DB
}

func NewDocumentTemplateRepository(db *gorm.DB) ports.IDocumentTemplateRepository {
	return &DocumentTemplateImpl{db: db}
}

// CreateDocumentTemplate implements ports.IDocumentTemplateRepository.
func (d *DocumentTemplateImpl) CreateDocumentTemplate(ctx context.Context, payload *models.DocumentTemplate) error {
	tx := database.HelperExtractTx(ctx, d.db)
	if err := tx.WithContext(ctx).Create(&payload).Error; err != nil {
		return err
	}
	return nil
}

// DeleteDocumentTemplate implements ports.IDocumentTemplateRepository.
func (d *DocumentTemplateImpl) DeleteDocumentTemplate(ctx context.Context, id uint) error {
	tx := database.HelperExtractTx(ctx, d.db)
	if err := tx.WithContext(ctx).Delete(&models.DocumentTemplate{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetDocumentTemplate implements ports.IDocumentTemplateRepository.
func (d *DocumentTemplateImpl) GetDocumentTemplate(ctx context.Context, id uint) (*models.DocumentTemplate, error) {
	tx := database.HelperExtractTx(ctx, d.db)
	var data models.DocumentTemplate
	if err := tx.WithContext(ctx).Where("id =?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// GetDocumentTemplates implements ports.IDocumentTemplateRepository.
func (d *DocumentTemplateImpl) GetDocumentTemplates(ctx context.Context) (*pagination.Pagination[[]models.DocumentTemplate], error) {
	tx := database.HelperExtractTx(ctx, d.db)
	p := pagination.GetFilters[filters.DocumentTemplateFilter](ctx)
	fp := p.Filters

	orderBy := pagination.NewOrderBy(pagination.SortParams{
		Sort:           p.Sort,
		Order:          p.Order,
		DefaultOrderBy: "updated_at DESC",
	})
	tx = pagination.ApplyFilter(tx, "id", fp.ID, "contains")
	tx = tx.WithContext(ctx).Order(orderBy)
	data, err := pagination.Paginate[filters.DocumentTemplateFilter, []models.DocumentTemplate](p, tx)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// UpdateDocumentTemplate implements ports.IDocumentTemplateRepository.
func (d *DocumentTemplateImpl) UpdateDocumentTemplate(ctx context.Context, payload *models.DocumentTemplate) error {
	tx := database.HelperExtractTx(ctx, d.db)
	if err := tx.WithContext(ctx).Save(&payload).Error; err != nil {
		return err
	}
	return nil
}
