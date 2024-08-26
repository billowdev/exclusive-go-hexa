package handlers

import (
	"context"
	"strconv"

	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	ports "github.com/billowdev/document-system-field-manager/internal/core/ports/system_fields"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/filters"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/pagination"
	"github.com/billowdev/document-system-field-manager/pkg/utils"
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
	SystemFieldImpls struct {
		systemFieldService ports.ISystemFieldService
	}
)

func NewSystemFieldHandler(
	systemFieldService ports.ISystemFieldService,
) ISystemFieldHandler {
	return &SystemFieldImpls{systemFieldService: systemFieldService}
}

// HandleCreateSystemField implements ISystemFieldHandler.
func (s *SystemFieldImpls) HandleCreateSystemField(c *fiber.Ctx) error {
	var payload models.SystemField
	if err := c.BodyParser(&payload); err != nil {
		return utils.NewErrorResponse(c, "Invalid request payload", err.Error())
	}

	res := s.systemFieldService.CreateSystemField(c.Context(), &payload)
	return c.JSON(res)
}

// HandleDeleteSystemField implements ISystemFieldHandler.
func (s *SystemFieldImpls) HandleDeleteSystemField(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid ID", err.Error())
	}

	res := s.systemFieldService.DeleteSystemField(c.Context(), uint(id))
	return c.JSON(res)
}

// HandleUpdateSystemField implements ISystemFieldHandler.
func (s *SystemFieldImpls) HandleUpdateSystemField(c *fiber.Ctx) error {
	var payload models.SystemField
	if err := c.BodyParser(&payload); err != nil {
		return utils.NewErrorResponse(c, "Invalid request payload", err.Error())
	}
	res := s.systemFieldService.UpdateSystemField(c.Context(), &payload)
	return c.JSON(res)
}

// HandleGetSystemField implements ISystemFieldHandler.
func (s *SystemFieldImpls) HandleGetSystemField(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid ID", err.Error())
	}

	res := s.systemFieldService.GetSystemField(c.Context(), uint(id))
	return c.JSON(res)
}

// HandleGetSystemFields implements ISystemFieldHandler.
func (s *SystemFieldImpls) HandleGetSystemFields(c *fiber.Ctx) error {
	ctx := context.Background()
	params := pagination.NewPaginationParams[filters.SystemFieldFilter](c)
	paramCtx := pagination.SetFilters(ctx, params)
	res := s.systemFieldService.GetSystemFields(paramCtx)
	return c.JSON(res)
}
