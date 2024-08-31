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

type LogDocumentVersionFieldValueImpl struct {
	db *gorm.DB
}

func NewLogDocumentVersionFieldValueRepository(db *gorm.DB) ports.ILogDocumentVersionFieldValueRepository {
	return &LogDocumentVersionFieldValueImpl{db: db}
}

// CreateLogDocumentVersionFieldValue implements ports.ILogDocumentVersionFieldValueRepository.
func (l *LogDocumentVersionFieldValueImpl) CreateLogDocumentVersionFieldValue(ctx context.Context, payload *models.LogDocumentVersionFieldValue) error {
	tx := database.HelperExtractTx(ctx, l.db)
	if err := tx.WithContext(ctx).Create(&payload).Error; err != nil {
		return err
	}
	return nil
}

// DeleteLogDocumentVersionFieldValue implements ports.ILogDocumentVersionFieldValueRepository.
func (l *LogDocumentVersionFieldValueImpl) DeleteLogDocumentVersionFieldValue(ctx context.Context, id uint) error {
	tx := database.HelperExtractTx(ctx, l.db)
	if err := tx.WithContext(ctx).Where("id=?", id).Delete(&models.LogDocumentVersionFieldValue{}).Error; err != nil {
		return err
	}
	return nil
}

// GetLogDocumentVersionFieldValue implements ports.ILogDocumentVersionFieldValueRepository.
func (l *LogDocumentVersionFieldValueImpl) GetLogDocumentVersionFieldValue(ctx context.Context, id uint) (*models.LogDocumentVersionFieldValue, error) {
	tx := database.HelperExtractTx(ctx, l.db)
	var data models.LogDocumentVersionFieldValue
	if err := tx.WithContext(ctx).Where("id =?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// GetLogDocumentVersionFieldValues implements ports.ILogDocumentVersionFieldValueRepository.
func (l *LogDocumentVersionFieldValueImpl) GetLogDocumentVersionFieldValues(ctx context.Context) (*pagination.Pagination[[]models.LogDocumentVersionFieldValue], error) {
	tx := database.HelperExtractTx(ctx, l.db)
	p := pagination.GetFilters[filters.LogDocumentVersionFieldValueFilter](ctx)
	fp := p.Filters

	orderBy := pagination.NewOrderBy(pagination.SortParams{
		Sort:           p.Sort,
		Order:          p.Order,
		DefaultOrderBy: "updated_at DESC",
	})
	tx = pagination.ApplyFilter(tx, "id", fp.ID, "contains")
	tx = tx.WithContext(ctx).Order(orderBy)
	data, err := pagination.Paginate[filters.LogDocumentVersionFieldValueFilter, []models.LogDocumentVersionFieldValue](p, tx)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// UpdateLogDocumentVersionFieldValue implements ports.ILogDocumentVersionFieldValueRepository.
func (l *LogDocumentVersionFieldValueImpl) UpdateLogDocumentVersionFieldValue(ctx context.Context, payload *models.LogDocumentVersionFieldValue) error {
	tx := database.HelperExtractTx(ctx, l.db)
	if err := tx.WithContext(ctx).Save(&payload).Error; err != nil {
		return err
	}
	return nil
}
