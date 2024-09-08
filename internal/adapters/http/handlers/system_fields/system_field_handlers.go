package handlers

import (
	"context"
	"strconv"
	"time"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	ports "github.com/billowdev/exclusive-go-hexa/internal/core/ports/system_fields"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/filters"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"github.com/billowdev/exclusive-go-hexa/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type (
	ISystemFieldHandler interface {
		HandleGetSystemField(c *fiber.Ctx) error
		HandleGetSystemFields(c *fiber.Ctx) error
		HandleUpdateSystemField(c *fiber.Ctx) error
		HandleCreateSystemField(c *fiber.Ctx) error
		HandleDeleteSystemField(c *fiber.Ctx) error
	}
	SystemFieldImpl struct {
		systemFieldService ports.ISystemFieldService
	}
)

func NewSystemFieldHandler(
	systemFieldService ports.ISystemFieldService,
) ISystemFieldHandler {
	return &SystemFieldImpl{systemFieldService: systemFieldService}
}

// HandleCreateSystemField implements ISystemFieldHandler.
func (s *SystemFieldImpl) HandleCreateSystemField(c *fiber.Ctx) error {
	var payload models.SystemField
	if err := c.BodyParser(&payload); err != nil {
		return utils.NewErrorResponse(c, "Invalid request payload", err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	res := s.systemFieldService.CreateSystemField(ctx, &payload)
	return c.JSON(res)
}

// HandleDeleteSystemField implements ISystemFieldHandler.
func (s *SystemFieldImpl) HandleDeleteSystemField(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid ID", err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	res := s.systemFieldService.DeleteSystemField(ctx, uint(id))
	return c.JSON(res)
}

// HandleUpdateSystemField implements ISystemFieldHandler.
func (s *SystemFieldImpl) HandleUpdateSystemField(c *fiber.Ctx) error {
	var payload models.SystemField
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
	res := s.systemFieldService.UpdateSystemField(ctx, uint(id), &payload)
	return c.JSON(res)
}

// HandleGetSystemField implements ISystemFieldHandler.
func (s *SystemFieldImpl) HandleGetSystemField(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid ID", err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	res := s.systemFieldService.GetSystemField(ctx, uint(id))
	return c.JSON(res)
}

// HandleGetSystemFields implements ISystemFieldHandler.
func (s *SystemFieldImpl) HandleGetSystemFields(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	params := pagination.NewPaginationParams[filters.SystemFieldFilter](c)
	paramCtx := pagination.SetFilters(ctx, params)
	res := s.systemFieldService.GetSystemFields(paramCtx)
	return c.JSON(res)
}
