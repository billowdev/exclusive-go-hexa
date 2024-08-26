package ports

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/filters"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"github.com/billowdev/exclusive-go-hexa/pkg/utils"
)

type IOrderRepository interface {
	GetOrder(ctx context.Context, id uint) (*models.Order, error)
	GetOrders(ctx context.Context, p pagination.PaginationParams[filters.OrderFilter]) (*pagination.Pagination[[]models.Order], error)
	CreateOrder(ctx context.Context, payload *models.Order) error
	UpdateOrder(ctx context.Context, payload *models.Order) error
	DeleteOrder(ctx context.Context, id uint) error
}

type IOrderService interface {
	GetOrder(ctx context.Context, id uint) utils.APIResponse
	GetOrders(ctx context.Context, p pagination.PaginationParams[filters.OrderFilter]) pagination.Pagination[[]models.Order]
	CreateOrder(ctx context.Context, payload *models.Order) utils.APIResponse
	UpdateOrder(ctx context.Context, payload *models.Order) utils.APIResponse
	DeleteOrder(ctx context.Context, id uint) utils.APIResponse
}
