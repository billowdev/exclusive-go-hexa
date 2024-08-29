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

type ConfigSystemMasterFileFieldImpl struct {
	db *gorm.DB
}

func NewConfigSystemMasterFileFieldRepository(db *gorm.DB) ports.IConfigSystemMasterFileFieldRepository {
	return &ConfigSystemMasterFileFieldImpl{db: db}
}

// CreateConfigSystemMasterFileField implements ports.IConfigSystemMasterFileFieldRepository.
func (c *ConfigSystemMasterFileFieldImpl) CreateConfigSystemMasterFileField(ctx context.Context, payload *models.ConfigSystemMasterFileField) error {
	tx := database.HelperExtractTx(ctx, c.db)
	if err := tx.WithContext(ctx).Create(&payload).Error; err != nil {
		return err
	}
	return nil
}

// DeleteConfigSystemMasterFileField implements ports.IConfigSystemMasterFileFieldRepository.
func (c *ConfigSystemMasterFileFieldImpl) DeleteConfigSystemMasterFileField(ctx context.Context, id uint) error {
	tx := database.HelperExtractTx(ctx, c.db)
	if err := tx.WithContext(ctx).Delete(&models.ConfigSystemMasterFileField{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetConfigSystemMasterFileField implements ports.IConfigSystemMasterFileFieldRepository.
func (c *ConfigSystemMasterFileFieldImpl) GetConfigSystemMasterFileField(ctx context.Context, id uint) (*models.ConfigSystemMasterFileField, error) {
	tx := database.HelperExtractTx(ctx, c.db)
	var data models.ConfigSystemMasterFileField
	if err := tx.WithContext(ctx).Where("id =?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// GetConfigSystemMasterFileFields implements ports.IConfigSystemMasterFileFieldRepository.
func (c *ConfigSystemMasterFileFieldImpl) GetConfigSystemMasterFileFields(ctx context.Context) (*pagination.Pagination[[]models.ConfigSystemMasterFileField], error) {
	tx := database.HelperExtractTx(ctx, c.db)
	p := pagination.GetFilters[filters.ConfigSystemMasterFileFieldFilter](ctx)
	fp := p.Filters

	orderBy := pagination.NewOrderBy(pagination.SortParams{
		Sort:           p.Sort,
		Order:          p.Order,
		DefaultOrderBy: "updated_at DESC",
	})
	tx = pagination.ApplyFilter(tx, "id", fp.ID, "contains")
	tx = tx.WithContext(ctx).Order(orderBy)
	data, err := pagination.Paginate[filters.ConfigSystemMasterFileFieldFilter, []models.ConfigSystemMasterFileField](p, tx)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// UpdateConfigSystemMasterFileField implements ports.IConfigSystemMasterFileFieldRepository.
func (c *ConfigSystemMasterFileFieldImpl) UpdateConfigSystemMasterFileField(ctx context.Context, payload *models.ConfigSystemMasterFileField) error {
	tx := database.HelperExtractTx(ctx, c.db)
	if err := tx.WithContext(ctx).Save(&payload).Error; err != nil {
		return err
	}
	return nil
}
