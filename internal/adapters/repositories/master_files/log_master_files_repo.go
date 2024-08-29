package repositories

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database"
	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	ports "github.com/billowdev/exclusive-go-hexa/internal/core/ports/master_files"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/filters"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"gorm.io/gorm"
)

type LogMasterFileRepository struct {
	db *gorm.DB
}

func NewLogMasterFileRepository(db *gorm.DB) ports.ILogMasterFileRepository {
	return LogMasterFileRepository{db: db}
}

// CreateLogMasterFile implements ports.ILogMasterFileRepository.
func (l LogMasterFileRepository) CreateLogMasterFile(ctx context.Context, payload *models.LogMasterFile) error {
	tx := database.HelperExtractTx(ctx, l.db)
	if err := tx.WithContext(ctx).Create(&payload).Error; err != nil {
		return err
	}
	return nil

}

// DeleteLogMasterFile implements ports.ILogMasterFileRepository.
func (l LogMasterFileRepository) DeleteLogMasterFile(ctx context.Context, id uint) error {
	tx := database.HelperExtractTx(ctx, l.db)
	if err := tx.WithContext(ctx).Delete(&models.LogMasterFile{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetLogMasterFile implements ports.ILogMasterFileRepository.
func (l LogMasterFileRepository) GetLogMasterFile(ctx context.Context, id uint) (*models.LogMasterFile, error) {
	tx := database.HelperExtractTx(ctx, l.db)
	var data models.LogMasterFile
	if err := tx.WithContext(ctx).Where("id =?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// GetLogMasterFiles implements ports.ILogMasterFileRepository.
func (l LogMasterFileRepository) GetLogMasterFiles(ctx context.Context) (*pagination.Pagination[[]models.LogMasterFile], error) {
	tx := database.HelperExtractTx(ctx, l.db)
	p := pagination.GetFilters[filters.LogMasterFileFilter](ctx)
	fp := p.Filters

	orderBy := pagination.NewOrderBy(pagination.SortParams{
		Sort:           p.Sort,
		Order:          p.Order,
		DefaultOrderBy: "updated_at DESC",
	})
	tx = pagination.ApplyFilter(tx, "id", fp.ID, "contains")
	tx = tx.WithContext(ctx).Order(orderBy)
	data, err := pagination.Paginate[filters.LogMasterFileFilter, []models.LogMasterFile](p, tx)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// UpdateLogMasterFile implements ports.ILogMasterFileRepository.
func (l LogMasterFileRepository) UpdateLogMasterFile(ctx context.Context, payload *models.LogMasterFile) error {
	tx := database.HelperExtractTx(ctx, l.db)
	if err := tx.WithContext(ctx).Save(&payload).Error; err != nil {
		return err
	}
	return nil
}
