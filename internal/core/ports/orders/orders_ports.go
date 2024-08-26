package ports

import (
	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/filters"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/pagination"
	"github.com/billowdev/document-system-field-manager/pkg/utils"
)

type IOrderRepository interface {
	GetOrder(id uint) (*models.Order, error)
	GetOrders(p pagination.PaginationParams[filters.OrderFilter]) (*pagination.Pagination[[]models.Order], error)
	CreateOrder(payload *models.Order) error
	UpdateOrder(payload *models.Order) error
	DeleteOrder(id uint) error
}

type IOrderService interface {
	GetOrder(id uint) utils.APIResponse
	GetOrders(p pagination.PaginationParams[filters.OrderFilter]) pagination.Pagination[[]models.Order]
	CreateOrder(payload *models.Order) utils.APIResponse
	UpdateOrder(payload *models.Order) utils.APIResponse
	DeleteOrder(id uint) utils.APIResponse
}
