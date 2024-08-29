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

type DocumentVersionImpl struct {
	db *gorm.DB
}

func NewDocumentVersionRepository(db *gorm.DB) ports.IDocumentVersionRepository {
	return &DocumentVersionImpl{db: db}
}

// CreateDocumentVersion implements ports.IDocumentVersionRepository.
func (d *DocumentVersionImpl) CreateDocumentVersion(ctx context.Context, payload *models.DocumentVersion) error {
	tx := database.HelperExtractTx(ctx, d.db)
	if err := tx.WithContext(ctx).Create(&payload).Error; err != nil {
		return err
	}
	return nil
}

// DeleteDocumentVersion implements ports.IDocumentVersionRepository.
func (d *DocumentVersionImpl) DeleteDocumentVersion(ctx context.Context, id uint) error {
	tx := database.HelperExtractTx(ctx, d.db)
	if err := tx.WithContext(ctx).Delete(&models.DocumentVersion{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetDocumentVersion implements ports.IDocumentVersionRepository.
func (d *DocumentVersionImpl) GetDocumentVersion(ctx context.Context, id uint) (*models.DocumentVersion, error) {
	tx := database.HelperExtractTx(ctx, d.db)
	var data models.DocumentVersion
	if err := tx.WithContext(ctx).Where("id =?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// GetDocumentVersions implements ports.IDocumentVersionRepository.
func (d *DocumentVersionImpl) GetDocumentVersions(ctx context.Context) (*pagination.Pagination[[]models.DocumentVersion], error) {
	tx := database.HelperExtractTx(ctx, d.db)
	p := pagination.GetFilters[filters.DocumentVersionFilter](ctx)
	fp := p.Filters

	orderBy := pagination.NewOrderBy(pagination.SortParams{
		Sort:           p.Sort,
		Order:          p.Order,
		DefaultOrderBy: "updated_at DESC",
	})
	tx = pagination.ApplyFilter(tx, "id", fp.ID, "contains")
	tx = tx.WithContext(ctx).Order(orderBy)
	data, err := pagination.Paginate[filters.DocumentVersionFilter, []models.DocumentVersion](p, tx)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// UpdateDocumentVersion implements ports.IDocumentVersionRepository.
func (d *DocumentVersionImpl) UpdateDocumentVersion(ctx context.Context, payload *models.DocumentVersion) error {
	tx := database.HelperExtractTx(ctx, d.db)
	if err := tx.WithContext(ctx).Save(&payload).Error; err != nil {
		return err
	}
	return nil
}
