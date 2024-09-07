package ports

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	domain "github.com/billowdev/exclusive-go-hexa/internal/core/domain/orders"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"github.com/billowdev/exclusive-go-hexa/pkg/utils"
)

type IOrderRepository interface {
	GetOrder(ctx context.Context, id uint) (*models.Order, error)
	GetOrders(ctx context.Context) (*pagination.Pagination[[]models.Order], error)
	CreateOrder(ctx context.Context, payload *models.Order) error
	UpdateOrder(ctx context.Context, payload *models.Order) error
	DeleteOrder(ctx context.Context, id uint) error
}

type IOrderService interface {
	GetOrder(ctx context.Context, id uint) utils.APIResponse
	GetOrders(ctx context.Context) pagination.Pagination[[]domain.OrderDomain]
	CreateOrder(ctx context.Context, payload domain.OrderDomain) utils.APIResponse
	UpdateOrder(ctx context.Context, id uint, payload domain.OrderDomain) utils.APIResponse
	DeleteOrder(ctx context.Context, id uint) utils.APIResponse
}
