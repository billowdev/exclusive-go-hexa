package domain

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database"
	domain "github.com/billowdev/exclusive-go-hexa/internal/core/domain/orders"
	ports "github.com/billowdev/exclusive-go-hexa/internal/core/ports/orders"
	"github.com/billowdev/exclusive-go-hexa/pkg/configs"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"github.com/billowdev/exclusive-go-hexa/pkg/utils"
)

type OrderServiceImpl struct {
	repo       ports.IOrderRepository
	transactor database.IDatabaseTransactor
}

func NewOrderService(
	repo ports.IOrderRepository,
	transactor database.IDatabaseTransactor,
) ports.IOrderService {
	return &OrderServiceImpl{repo: repo, transactor: transactor}
}

// CreateOrder implements ports.IOrderService.
func (o *OrderServiceImpl) CreateOrder(ctx context.Context, payload domain.OrderDomain) utils.APIResponse {
	data := domain.ToOrderModel(payload)
	if err := o.repo.CreateOrder(ctx, data); err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: nil}
}

// DeleteOrder implements ports.IOrderService.
func (o *OrderServiceImpl) DeleteOrder(ctx context.Context, id uint) utils.APIResponse {
	if err := o.repo.DeleteOrder(ctx, id); err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: nil}
}

// GetOrder implements ports.IOrderService.
func (o *OrderServiceImpl) GetOrder(ctx context.Context, id uint) utils.APIResponse {
	data, err := o.repo.GetOrder(ctx, id)
	if err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	if data == nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Not Found", Data: nil}
	}
	res := domain.ToOrderDomain(data)
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: res}
}

// GetOrders implements ports.IOrderService.
func (o *OrderServiceImpl) GetOrders(ctx context.Context) pagination.Pagination[[]domain.OrderDomain] {
	data, err := o.repo.GetOrders(ctx)
	if err != nil {
		return pagination.Pagination[[]domain.OrderDomain]{}
	}
	// Convert repository data to domain models
	newData := utils.ConvertSlice(data.Rows, domain.ToOrderDomain)
	return pagination.Pagination[[]domain.OrderDomain]{
		Rows:       newData,
		Links:      data.Links,
		Total:      data.Total,
		Page:       data.Page,
		PageSize:   data.PageSize,
		TotalPages: data.TotalPages,
	}
}

// UpdateOrder implements ports.IOrderService.
func (o *OrderServiceImpl) UpdateOrder(ctx context.Context, id uint, payload domain.OrderDomain) utils.APIResponse {
	if _, err := o.repo.GetOrder(ctx, id); err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Not Found", Data: nil}
	}
	payload.ID = id
	data := domain.ToOrderModel(payload)
	if err := o.repo.UpdateOrder(ctx, data); err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	res := domain.ToOrderDomain(data)
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: res}
}
