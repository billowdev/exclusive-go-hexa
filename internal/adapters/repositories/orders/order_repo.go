package repositories

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database"
	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	ports "github.com/billowdev/exclusive-go-hexa/internal/core/ports/orders"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/filters"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"gorm.io/gorm"
)

type OrderImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) ports.IOrderRepository {
	return &OrderImpl{db: db}
}

// CreateOrder implements ports.IOrderRepository.
func (o *OrderImpl) CreateOrder(ctx context.Context, payload *models.Order) error {
	tx := database.HelperExtractTx(ctx, o.db)
	if err := tx.WithContext(ctx).Create(&payload).Error; err != nil {
		return err
	}
	return nil
}

// DeleteOrder implements ports.IOrderRepository.
func (o *OrderImpl) DeleteOrder(ctx context.Context, id uint) error {
	tx := database.HelperExtractTx(ctx, o.db)
	if err := tx.WithContext(ctx).Where("id=?", id).Delete(&models.Order{}).Error; err != nil {
		return err
	}
	return nil
}

// GetOrder implements ports.IOrderRepository.
func (o *OrderImpl) GetOrder(ctx context.Context, id uint) (*models.Order, error) {
	tx := database.HelperExtractTx(ctx, o.db)

	var data models.Order
	if err := tx.WithContext(ctx).Where("id =?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// GetOrders implements ports.IOrderRepository.
func (o *OrderImpl) GetOrders(ctx context.Context) (*pagination.Pagination[[]models.Order], error) {
	tx := database.HelperExtractTx(ctx, o.db)

	p := pagination.GetFilters[filters.OrderFilter](ctx)
	fp := p.Filters

	orderBy := pagination.NewOrderBy(pagination.SortParams{
		Sort:           p.Sort,
		Order:          p.Order,
		DefaultOrderBy: "updated_at DESC",
	})
	tx = pagination.ApplyFilter(tx, "id", fp.ID, "contains")
	tx = tx.WithContext(ctx).Order(orderBy)
	data, err := pagination.Paginate[filters.OrderFilter, []models.Order](p, tx)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// UpdateOrder implements ports.IOrderRepository.
func (o *OrderImpl) UpdateOrder(ctx context.Context, payload *models.Order) error {
	tx := database.HelperExtractTx(ctx, o.db)
	if err := tx.WithContext(ctx).Save(&payload).Error; err != nil {
		return err
	}
	return nil
}
