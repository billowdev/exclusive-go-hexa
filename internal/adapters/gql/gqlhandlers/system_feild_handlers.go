package gqlhandlers

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	"github.com/billowdev/exclusive-go-hexa/internal/adapters/gql/model"
	ports "github.com/billowdev/exclusive-go-hexa/internal/core/ports/system_fields"
)

type (
	IGQLSystemFieldHandler interface {
		CreateSystemField(req *model.SystemField) *model.SystemField
	}
	GQLSystemFieldHandlerImpl struct {
		systemFieldService ports.ISystemFieldService // dependency injection
	}
)

func NewGQLSystemFieldHandler(systemFieldService ports.ISystemFieldService) IGQLSystemFieldHandler {
	return &GQLSystemFieldHandlerImpl{systemFieldService: systemFieldService}
}

// CreateSystemField implements IGQLSystemFieldHandler.
func (g *GQLSystemFieldHandlerImpl) CreateSystemField(req *model.SystemField) *model.SystemField {
	descripton := ""
	defaultValue := ""
	if req.Description != nil {
		descripton = *req.Description
	}
	if req.DefaultValue != nil {
		defaultValue = *req.DefaultValue
	}
	systemField := &models.SystemField{
		FieldCode:    req.FieldCode,
		FieldName:    req.FieldName,
		DataType:     req.DataType,
		Description:  descripton,
		DefaultValue: defaultValue,
	}
	response := g.systemFieldService.CreateSystemField(context.TODO(), systemField)
	_ = response
	return req
}
