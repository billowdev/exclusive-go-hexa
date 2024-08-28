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

type SystemGroupFieldRepositoryImpls struct {
	db *gorm.DB
}

func NewSystemGroupFieldRepository(db *gorm.DB) ports.ISystemGroupFieldRepository {
	return &SystemGroupFieldRepositoryImpls{db: db}
}

// CreateSystemGroupField implements ports.ISystemGroupFieldRepository.
func (s *SystemGroupFieldRepositoryImpls) CreateSystemGroupField(ctx context.Context, payload *models.SystemGroupField) error {
	tx := database.ExtractTx(ctx)
	if tx == nil {
		tx = s.db
	}
	if err := tx.Create(&payload).Error; err != nil {
		return err
	}
	return nil
}

// DeleteSystemGroupField implements ports.ISystemGroupFieldRepository.
func (s *SystemGroupFieldRepositoryImpls) DeleteSystemGroupField(ctx context.Context, id uint) error {
	tx := database.ExtractTx(ctx)
	if tx == nil {
		tx = s.db
	}
	if err := tx.Delete(&models.SystemGroupField{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetSystemGroupField implements ports.ISystemGroupFieldRepository.
func (s *SystemGroupFieldRepositoryImpls) GetSystemGroupField(ctx context.Context, id uint) (*models.SystemGroupField, error) {
	tx := database.ExtractTx(ctx)
	if tx == nil {
		tx = s.db
	}
	var data models.SystemGroupField
	if err := tx.Where("id =?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// GetSystemGroupFields implements ports.ISystemGroupFieldRepository.
func (s *SystemGroupFieldRepositoryImpls) GetSystemGroupFields(ctx context.Context) (*pagination.Pagination[[]models.SystemGroupField], error) {
	p := pagination.GetFilters[filters.SystemGroupFieldFilter](ctx)
	fp := p.Filters
	tx := database.ExtractTx(ctx)
	if tx == nil {
		tx = s.db
	}
	orderBy := pagination.NewOrderBy(pagination.SortParams{
		Sort:           p.Sort,
		Order:          p.Order,
		DefaultOrderBy: "updated_at DESC",
	})
	tx = pagination.ApplyFilter(tx, "id", fp.ID, "contains")
	tx = tx.Order(orderBy)
	pgR, err := pagination.Paginate[filters.SystemGroupFieldFilter, []models.SystemGroupField](p, tx)
	if err != nil {
		return nil, err
	}
	return &pgR, nil
}

// UpdateSystemGroupField implements ports.ISystemGroupFieldRepository.
func (s *SystemGroupFieldRepositoryImpls) UpdateSystemGroupField(ctx context.Context, payload *models.SystemGroupField) error {
	tx := database.ExtractTx(ctx)
	if tx == nil {
		tx = s.db
	}
	if err := tx.Save(&payload).Error; err != nil {
		return err
	}
	return nil
}
