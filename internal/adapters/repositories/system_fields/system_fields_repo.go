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

type SystemFieldRepositoryImpl struct {
	db *gorm.DB
}

func NewSystemFieldRepo(db *gorm.DB) ports.ISystemFieldRepository {
	return &SystemFieldRepositoryImpl{db: db}
}

// CreateSystemField implements ports.ISystemFieldRepository.
func (s *SystemFieldRepositoryImpl) CreateSystemField(ctx context.Context, payload *models.SystemField) error {
	tx := database.HelperExtractTx(ctx, s.db)
	if err := tx.WithContext(ctx).Create(&payload).Error; err != nil {
		return err
	}
	return nil
}

// DeleteSystemField implements ports.ISystemFieldRepository.
func (s *SystemFieldRepositoryImpl) DeleteSystemField(ctx context.Context, id uint) error {
	tx := database.HelperExtractTx(ctx, s.db)
	if err := tx.WithContext(ctx).Delete(&models.SystemField{}, id).Error; err != nil {
		return err
	}
	return nil
	// TODO: Handle soft deletion or other forms of soft deletion. For example, update a deleted_at field instead of deleting the record directly.
	// return s.db.Exec("UPDATE system_fields SET deleted_at = ? WHERE id =?", time.Now(), id).Error
	// return s.db.Delete(&models.SystemField{}, id).Error
	// return s.db.Model(&models.SystemField{}).Where("id =?", id).Update("deleted_at", time.Now()).Error
	// return s.db.Where("id =?", id).Delete(&models.SystemField{}).Error
	// return s.db.Exec("UPDATE system_fields SET deleted_at = CURRENT_TIMESTAMP WHERE id =?", id).
}

// GetSystemField implements ports.ISystemFieldRepository.
func (s *SystemFieldRepositoryImpl) GetSystemField(ctx context.Context, id uint) (*models.SystemField, error) {
	tx := database.HelperExtractTx(ctx, s.db)

	var systemField models.SystemField
	if err := tx.WithContext(ctx).Where("id =?", id).First(&systemField).Error; err != nil {
		return nil, err
	}
	return &systemField, nil
	// TODO: Handle soft deleted records. For example, check if deleted_at is not NULL before returning the record.
	// var systemField models.SystemField
	// if err := s.db.Where("id =? AND deleted_at IS NULL", id).First(&systemField).Error; err!= nil {
	//     return nil, err
	// }
	// return &systemField, nil
	// return s.db.Where("id =?", id).First(&models.SystemField{}).Error
	// return s.db.Where("id =?", id).First(&systemField).Error
	// return s.db.Exec("
}

// GetSystemFields implements ports.ISystemFieldRepository.
func (s *SystemFieldRepositoryImpl) GetSystemFields(ctx context.Context) (*pagination.Pagination[[]models.SystemField], error) {
	p := pagination.GetFilters[filters.SystemFieldFilter](ctx)
	fp := p.Filters
	q := s.db
	orderBy := pagination.NewOrderBy(pagination.SortParams{
		Sort:           p.Sort,
		Order:          p.Order,
		DefaultOrderBy: "updated_at DESC",
	})
	q = pagination.ApplyFilter(q, "id", fp.ID, "contains")
	q = q.Order(orderBy)
	pgR, err := pagination.Paginate[filters.SystemFieldFilter, []models.SystemField](p, q)
	if err != nil {
		return nil, err
	}
	return &pgR, nil
}

// UpdateSystemField implements ports.ISystemFieldRepository.
func (s *SystemFieldRepositoryImpl) UpdateSystemField(ctx context.Context, payload *models.SystemField) error {
	tx := database.HelperExtractTx(ctx, s.db)

	if err := tx.WithContext(ctx).Save(&payload).Error; err != nil {
		return err
	}
	return nil
	// TODO: Handle soft deletion or other forms of soft deletion. For example, update a deleted_at field instead of deleting the record directly.
	// return s.db.Exec("UPDATE system_fields SET deleted_at = NULL WHERE id =?", payload.ID).Error
	// return s.db.Save(&payload).Error
	// return s.db.Where("id =?", payload.ID).Updates(payload).Error
	// return s.db.Exec("
}
