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

type DocumentImpl struct {
	db *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) ports.IDocumentRepository {
	return &DocumentImpl{db: db}
}

// CreateDocument implements ports.IDocumentRepository.
func (d *DocumentImpl) CreateDocument(ctx context.Context, payload *models.Document) error {
	tx := database.HelperExtractTx(ctx, d.db)
	if err := tx.WithContext(ctx).Create(&payload).Error; err != nil {
		return err
	}
	return nil
}

// DeleteDocument implements ports.IDocumentRepository.
func (d *DocumentImpl) DeleteDocument(ctx context.Context, id uint) error {
	tx := database.HelperExtractTx(ctx, d.db)
	if err := tx.WithContext(ctx).Delete(&models.Document{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetDocument implements ports.IDocumentRepository.
func (d *DocumentImpl) GetDocument(ctx context.Context, id uint) (*models.Document, error) {
	tx := database.HelperExtractTx(ctx, d.db)
	var data models.Document
	if err := tx.WithContext(ctx).Where("id =?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// GetDocuments implements ports.IDocumentRepository.
func (d *DocumentImpl) GetDocuments(ctx context.Context) (*pagination.Pagination[[]models.Document], error) {
	tx := database.HelperExtractTx(ctx, d.db)
	p := pagination.GetFilters[filters.DocumentFilter](ctx)
	fp := p.Filters

	orderBy := pagination.NewOrderBy(pagination.SortParams{
		Sort:           p.Sort,
		Order:          p.Order,
		DefaultOrderBy: "updated_at DESC",
	})
	tx = pagination.ApplyFilter(tx, "id", fp.ID, "contains")
	tx = tx.WithContext(ctx).Order(orderBy)
	data, err := pagination.Paginate[filters.DocumentFilter, []models.Document](p, tx)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// UpdateDocument implements ports.IDocumentRepository.
func (d *DocumentImpl) UpdateDocument(ctx context.Context, payload *models.Document) error {
	tx := database.HelperExtractTx(ctx, d.db)

	if err := tx.WithContext(ctx).Save(&payload).Error; err != nil {
		return err
	}
	return nil
}
