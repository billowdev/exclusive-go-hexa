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

type MasterFileImpl struct {
	db *gorm.DB
}

func NewMasterFileRepository(db *gorm.DB) ports.IMasterFileRepository {
	return &MasterFileImpl{db: db}
}

// CreateMasterFile implements ports.IMasterFileRepository.
func (m *MasterFileImpl) CreateMasterFile(ctx context.Context, payload *models.MasterFile) error {
	tx := database.HelperExtractTx(ctx, m.db)
	if err := tx.WithContext(ctx).Create(&payload).Error; err != nil {
		return err
	}
	return nil
}

// DeleteMasterFile implements ports.IMasterFileRepository.
func (m *MasterFileImpl) DeleteMasterFile(ctx context.Context, id uint) error {
	tx := database.HelperExtractTx(ctx, m.db)
	if err := tx.WithContext(ctx).Where("id=?", id).Delete(&models.MasterFile{}).Error; err != nil {
		return err
	}
	return nil
}

// GetMasterFile implements ports.IMasterFileRepository.
func (m *MasterFileImpl) GetMasterFile(ctx context.Context, id uint) (*models.MasterFile, error) {
	tx := database.HelperExtractTx(ctx, m.db)
	var data models.MasterFile
	if err := tx.WithContext(ctx).Where("id =?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// GetMasterFiles implements ports.IMasterFileRepository.
func (m *MasterFileImpl) GetMasterFiles(ctx context.Context) (*pagination.Pagination[[]models.MasterFile], error) {
	tx := database.HelperExtractTx(ctx, m.db)
	p := pagination.GetFilters[filters.MasterFileFilter](ctx)
	fp := p.Filters

	orderBy := pagination.NewOrderBy(pagination.SortParams{
		Sort:           p.Sort,
		Order:          p.Order,
		DefaultOrderBy: "updated_at DESC",
	})
	tx = pagination.ApplyFilter(tx, "id", fp.ID, "contains")
	tx = tx.WithContext(ctx).Order(orderBy)
	data, err := pagination.Paginate[filters.MasterFileFilter, []models.MasterFile](p, tx)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// UpdateMasterFile implements ports.IMasterFileRepository.
func (m *MasterFileImpl) UpdateMasterFile(ctx context.Context, payload *models.MasterFile) error {
	tx := database.HelperExtractTx(ctx, m.db)
	if err := tx.WithContext(ctx).Save(&payload).Error; err != nil {
		return err
	}
	return nil
}
