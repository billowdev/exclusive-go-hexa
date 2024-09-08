package handlers

import (
	"context"
	"strconv"
	"time"

	domain "github.com/billowdev/exclusive-go-hexa/internal/core/domain/orders"
	ports "github.com/billowdev/exclusive-go-hexa/internal/core/ports/orders"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/filters"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"github.com/billowdev/exclusive-go-hexa/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type (
	IOrderHandler interface {
		HandleGetOrder(c *fiber.Ctx) error
		HandleGetOrders(c *fiber.Ctx) error
		HandleUpdateOrder(c *fiber.Ctx) error
		HandleCreateOrder(c *fiber.Ctx) error
		HandleDeleteOrder(c *fiber.Ctx) error
	}
	OrderImpl struct {
		orderService ports.IOrderService
	}
)

func NewOrderHandler(
	orderService ports.IOrderService,
) IOrderHandler {
	return &OrderImpl{
		orderService: orderService,
	}
}

// HandleCreateOrder implements IOrderHandler.
func (h *OrderImpl) HandleCreateOrder(c *fiber.Ctx) error {
	var payload domain.OrderDomain
	if err := c.BodyParser(&payload); err != nil {
		return utils.NewErrorResponse(c, "Invalid request payload", err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	res := h.orderService.CreateOrder(ctx, payload)
	return c.JSON(res)
}

// HandleDeleteOrder implements IOrderHandler.
func (h *OrderImpl) HandleDeleteOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid ID", err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	res := h.orderService.DeleteOrder(ctx, uint(id))
	return c.JSON(res)
}

// HandleUpdateOrder implements IOrderHandler.
func (h *OrderImpl) HandleUpdateOrder(c *fiber.Ctx) error {
	var payload domain.OrderDomain
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid ID", err.Error())
	}
	if err := c.BodyParser(&payload); err != nil {
		return utils.NewErrorResponse(c, "Invalid request payload", err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	res := h.orderService.UpdateOrder(ctx, uint(id), payload)
	return c.JSON(res)
}

// HandleGetOrder implements IOrderHandler.
func (h *OrderImpl) HandleGetOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid ID", err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	res := h.orderService.GetOrder(ctx, uint(id))
	return c.JSON(res)
}

// HandleGetOrders implements IOrderHandler.
func (h *OrderImpl) HandleGetOrders(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	params := pagination.NewPaginationParams[filters.OrderFilter](c)
	paramCtx := pagination.SetFilters(ctx, params)
	res := h.orderService.GetOrders(paramCtx)
	return c.JSON(res)
}
