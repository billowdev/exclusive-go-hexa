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

type OrderImpls struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) ports.IOrderRepository {
	return &OrderImpls{db: db}
}

// CreateOrder implements ports.IOrderRepository.
func (o *OrderImpls) CreateOrder(ctx context.Context, payload *models.Order) error {
	tx := database.ExtractTx(ctx)
	if tx == nil {
		tx = o.db
	}
	if err := tx.Create(&payload).Error; err != nil {
		return err
	}
	return nil
}

// DeleteOrder implements ports.IOrderRepository.
func (o *OrderImpls) DeleteOrder(ctx context.Context, id uint) error {
	tx := database.ExtractTx(ctx)
	if tx == nil {
		tx = o.db
	}
	if err := tx.Delete(&models.Order{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetOrder implements ports.IOrderRepository.
func (o *OrderImpls) GetOrder(ctx context.Context, id uint) (*models.Order, error) {
	tx := database.ExtractTx(ctx)
	if tx == nil {
		tx = o.db
	}
	var data models.Order
	if err := tx.Where("id =?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// GetOrders implements ports.IOrderRepository.
func (o *OrderImpls) GetOrders(ctx context.Context) (*pagination.Pagination[[]models.Order], error) {
	tx := database.ExtractTx(ctx)
	if tx == nil {
		tx = o.db
	}
	p := pagination.GetFilters[filters.OrderFilter](ctx)
	fp := p.Filters

	orderBy := pagination.NewOrderBy(pagination.SortParams{
		Sort:           p.Sort,
		Order:          p.Order,
		DefaultOrderBy: "updated_at DESC",
	})
	tx = pagination.ApplyFilter(tx, "id", fp.ID, "contains")
	tx = tx.Order(orderBy)
	data, err := pagination.Paginate[filters.OrderFilter, []models.Order](p, tx)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// UpdateOrder implements ports.IOrderRepository.
func (o *OrderImpls) UpdateOrder(ctx context.Context, payload *models.Order) error {
	panic("unimplemented")
}
