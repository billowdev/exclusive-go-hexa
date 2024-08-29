package repositories

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database"
	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	ports "github.com/billowdev/exclusive-go-hexa/internal/core/ports/activities"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"gorm.io/gorm"
)

type ActivityRepositoryImpl struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) ports.IActivityRepository {
	return &ActivityRepositoryImpl{db: db}
}

// CreateActivity implements ports.IActivityRepository.
func (a *ActivityRepositoryImpl) CreateActivity(ctx context.Context, payload *models.Activity) error {
	tx := database.HelperExtractTx(ctx, a.db)
	if err := tx.WithContext(ctx).Create(&payload).Error; err != nil {
		return err
	}
	return nil
}

// DeleteActivity implements ports.IActivityRepository.
func (a *ActivityRepositoryImpl) DeleteActivity(ctx context.Context, id uint) error {
	tx := database.HelperExtractTx(ctx, a.db)
	if err := tx.WithContext(ctx).Delete(&models.Activity{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetActivities implements ports.IActivityRepository.
func (a *ActivityRepositoryImpl) GetActivities(ctx context.Context) (*pagination.Pagination[[]models.Activity], error) {
	tx := database.HelperExtractTx(ctx, a.db)
	p := pagination.GetFilters[models.Activity](ctx)
	fp := p.Filters

	orderBy := pagination.NewOrderBy(pagination.SortParams{
		Sort:           p.Sort,
		Order:          p.Order,
		DefaultOrderBy: "updated_at DESC",
	})
	tx = pagination.ApplyFilter(tx, "id", fp.ID, "contains")
	tx = tx.WithContext(ctx).Order(orderBy)
	data, err := pagination.Paginate[models.Activity, []models.Activity](p, tx)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetActivity implements ports.IActivityRepository.
func (a *ActivityRepositoryImpl) GetActivity(ctx context.Context, id uint) (*models.Activity, error) {
	tx := database.HelperExtractTx(ctx, a.db)
	var data models.Activity
	if err := tx.WithContext(ctx).Where("id =?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// UpdateActivity implements ports.IActivityRepository.
func (a *ActivityRepositoryImpl) UpdateActivity(ctx context.Context, payload *models.Activity) error {
	tx := database.HelperExtractTx(ctx, a.db)
	if err := tx.WithContext(ctx).Save(&payload).Error; err != nil {
		return err
	}
	return nil
}
