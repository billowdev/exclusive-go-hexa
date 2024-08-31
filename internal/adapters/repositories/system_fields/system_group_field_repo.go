package repositories

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database"
	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	ports "github.com/billowdev/exclusive-go-hexa/internal/core/ports/system_fields"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/filters"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"gorm.io/gorm"
)

type SystemGroupFieldRepositoryImpl struct {
	db *gorm.DB
}

func NewSystemGroupFieldRepository(db *gorm.DB) ports.ISystemGroupFieldRepository {
	return &SystemGroupFieldRepositoryImpl{db: db}
}

// CreateSystemGroupField implements ports.ISystemGroupFieldRepository.
func (s *SystemGroupFieldRepositoryImpl) CreateSystemGroupField(ctx context.Context, payload *models.SystemGroupField) error {
	tx := database.HelperExtractTx(ctx, s.db)
	if err := tx.WithContext(ctx).Create(&payload).Error; err != nil {
		return err
	}
	return nil
}

// DeleteSystemGroupField implements ports.ISystemGroupFieldRepository.
func (s *SystemGroupFieldRepositoryImpl) DeleteSystemGroupField(ctx context.Context, id uint) error {
	tx := database.HelperExtractTx(ctx, s.db)
	if err := tx.WithContext(ctx).Where("id=?", id).Delete(&models.SystemGroupField{}).Error; err != nil {
		return err
	}
	return nil
}

// GetSystemGroupField implements ports.ISystemGroupFieldRepository.
func (s *SystemGroupFieldRepositoryImpl) GetSystemGroupField(ctx context.Context, id uint) (*models.SystemGroupField, error) {
	tx := database.HelperExtractTx(ctx, s.db)
	var data models.SystemGroupField
	if err := tx.WithContext(ctx).Where("id =?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// GetSystemGroupFields implements ports.ISystemGroupFieldRepository.
func (s *SystemGroupFieldRepositoryImpl) GetSystemGroupFields(ctx context.Context) (*pagination.Pagination[[]models.SystemGroupField], error) {
	p := pagination.GetFilters[filters.SystemGroupFieldFilter](ctx)
	fp := p.Filters
	tx := database.HelperExtractTx(ctx, s.db)
	orderBy := pagination.NewOrderBy(pagination.SortParams{
		Sort:           p.Sort,
		Order:          p.Order,
		DefaultOrderBy: "updated_at DESC",
	})
	tx = pagination.ApplyFilter(tx, "id", fp.ID, "contains")
	tx = tx.WithContext(ctx).Order(orderBy)
	pgR, err := pagination.Paginate[filters.SystemGroupFieldFilter, []models.SystemGroupField](p, tx)
	if err != nil {
		return nil, err
	}
	return &pgR, nil
}

// UpdateSystemGroupField implements ports.ISystemGroupFieldRepository.
func (s *SystemGroupFieldRepositoryImpl) UpdateSystemGroupField(ctx context.Context, payload *models.SystemGroupField) error {
	tx := database.HelperExtractTx(ctx, s.db)
	if err := tx.WithContext(ctx).Save(&payload).Error; err != nil {
		return err
	}
	return nil
}
