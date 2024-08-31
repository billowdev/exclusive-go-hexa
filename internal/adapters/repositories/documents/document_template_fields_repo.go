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

type DocumentTemplateFieldImpl struct {
	db *gorm.DB
}

func NewDocumentTemplateFieldRepository(db *gorm.DB) ports.IDocumentTemplateFieldRepository {
	return &DocumentTemplateFieldImpl{db: db}
}

// CreateDocumentTemplateField implements ports.IDocumentTemplateFieldRepository.
func (d *DocumentTemplateFieldImpl) CreateDocumentTemplateField(ctx context.Context, payload *models.DocumentTemplateField) error {
	tx := database.HelperExtractTx(ctx, d.db)
	if err := tx.WithContext(ctx).Create(&payload).Error; err != nil {
		return err
	}
	return nil
}

// DeleteDocumentTemplateField implements ports.IDocumentTemplateFieldRepository.
func (d *DocumentTemplateFieldImpl) DeleteDocumentTemplateField(ctx context.Context, id uint) error {
	tx := database.HelperExtractTx(ctx, d.db)
	if err := tx.WithContext(ctx).Where("id=?", id).Delete(&models.DocumentTemplateField{}).Error; err != nil {
		return err
	}
	return nil
}

// GetDocumentTemplateField implements ports.IDocumentTemplateFieldRepository.
func (d *DocumentTemplateFieldImpl) GetDocumentTemplateField(ctx context.Context, id uint) (*models.DocumentTemplateField, error) {
	tx := database.HelperExtractTx(ctx, d.db)
	var data models.DocumentTemplateField
	if err := tx.WithContext(ctx).Where("id =?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// GetDocumentTemplateFields implements ports.IDocumentTemplateFieldRepository.
func (d *DocumentTemplateFieldImpl) GetDocumentTemplateFields(ctx context.Context) (*pagination.Pagination[[]models.DocumentTemplateField], error) {
	tx := database.HelperExtractTx(ctx, d.db)
	p := pagination.GetFilters[filters.DocumentTemplateFieldFilter](ctx)
	fp := p.Filters

	orderBy := pagination.NewOrderBy(pagination.SortParams{
		Sort:           p.Sort,
		Order:          p.Order,
		DefaultOrderBy: "updated_at DESC",
	})
	tx = pagination.ApplyFilter(tx, "id", fp.ID, "contains")
	tx = tx.WithContext(ctx).Order(orderBy)
	data, err := pagination.Paginate[filters.DocumentTemplateFieldFilter, []models.DocumentTemplateField](p, tx)
	if err != nil {
		return nil, err
	}
	return &data, nil

}

// UpdateDocumentTemplateField implements ports.IDocumentTemplateFieldRepository.
func (d *DocumentTemplateFieldImpl) UpdateDocumentTemplateField(ctx context.Context, payload *models.DocumentTemplateField) error {
	tx := database.HelperExtractTx(ctx, d.db)
	if err := tx.WithContext(ctx).Save(&payload).Error; err != nil {
		return err
	}
	return nil
}
