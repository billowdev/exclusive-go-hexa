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

type DocumentVersionFieldValueImpl struct {
	db *gorm.DB
}

func NewDocumentVersionFieldValueRepository(db *gorm.DB) ports.IDocumentVersionFieldValueRepository {
	return &DocumentVersionFieldValueImpl{db: db}
}

// CreateDocumentVersionFieldValue implements ports.IDocumentVersionFieldValueRepository.
func (d *DocumentVersionFieldValueImpl) CreateDocumentVersionFieldValue(ctx context.Context, payload *models.DocumentVersionFieldValue) error {
	tx := database.HelperExtractTx(ctx, d.db)
	if err := tx.WithContext(ctx).Create(&payload).Error; err != nil {
		return err
	}
	return nil
}

// DeleteDocumentVersionFieldValue implements ports.IDocumentVersionFieldValueRepository.
func (d *DocumentVersionFieldValueImpl) DeleteDocumentVersionFieldValue(ctx context.Context, id uint) error {
	tx := database.HelperExtractTx(ctx, d.db)
	if err := tx.WithContext(ctx).Delete(&models.DocumentVersionFieldValue{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetDocumentVersionFieldValue implements ports.IDocumentVersionFieldValueRepository.
func (d *DocumentVersionFieldValueImpl) GetDocumentVersionFieldValue(ctx context.Context, id uint) (*models.DocumentVersionFieldValue, error) {
	tx := database.HelperExtractTx(ctx, d.db)
	var data models.DocumentVersionFieldValue
	if err := tx.WithContext(ctx).Where("id =?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// GetDocumentVersionFieldValues implements ports.IDocumentVersionFieldValueRepository.
func (d *DocumentVersionFieldValueImpl) GetDocumentVersionFieldValues(ctx context.Context) (*pagination.Pagination[[]models.DocumentVersionFieldValue], error) {
	tx := database.HelperExtractTx(ctx, d.db)
	p := pagination.GetFilters[filters.DocumentVersionFieldValueFilter](ctx)
	fp := p.Filters

	orderBy := pagination.NewOrderBy(pagination.SortParams{
		Sort:           p.Sort,
		Order:          p.Order,
		DefaultOrderBy: "updated_at DESC",
	})
	tx = pagination.ApplyFilter(tx, "id", fp.ID, "contains")
	tx = tx.WithContext(ctx).Order(orderBy)
	data, err := pagination.Paginate[filters.DocumentVersionFieldValueFilter, []models.DocumentVersionFieldValue](p, tx)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// UpdateDocumentVersionFieldValue implements ports.IDocumentVersionFieldValueRepository.
func (d *DocumentVersionFieldValueImpl) UpdateDocumentVersionFieldValue(ctx context.Context, payload *models.DocumentVersionFieldValue) error {
	tx := database.HelperExtractTx(ctx, d.db)
	if err := tx.WithContext(ctx).Save(&payload).Error; err != nil {
		return err
	}
	return nil
}
