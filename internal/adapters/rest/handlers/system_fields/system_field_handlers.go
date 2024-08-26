package handlers

import (
	"context"

	ports "github.com/billowdev/document-system-field-manager/internal/core/ports/system_fields"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/filters"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/pagination"
	"github.com/gofiber/fiber/v2"
)

type (
	ISystemFieldHandler interface {
		HandleGetSystemField(c *fiber.Ctx) error
		HandleGetSystemFields(c *fiber.Ctx) error
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

// HandleGetSystemField implements ISystemFieldHandler.
func (s *SystemFieldImpls) HandleGetSystemField(c *fiber.Ctx) error {
	panic("unimplemented")
}

// HandleGetSystemFields implements ISystemFieldHandler.
func (s *SystemFieldImpls) HandleGetSystemFields(c *fiber.Ctx) error {
	ctx := context.Background()
	params := pagination.NewPaginationParams[filters.SystemFieldFilter](c)

	paramCtx := pagination.SetFilters[filters.SystemFieldFilter](ctx, params)
	res := s.systemFieldService.GetSystemFields(paramCtx)
	return c.JSON(res)
}
