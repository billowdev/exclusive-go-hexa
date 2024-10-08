package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database"
	graph "github.com/billowdev/exclusive-go-hexa/internal/adapters/gql"
	"github.com/billowdev/exclusive-go-hexa/internal/adapters/gql/gqlhandlers"
	"github.com/billowdev/exclusive-go-hexa/internal/adapters/gql/model"
	repositories "github.com/billowdev/exclusive-go-hexa/internal/adapters/repositories/system_fields"
	services "github.com/billowdev/exclusive-go-hexa/internal/core/services/system_fields"
	"gorm.io/gorm"
)

// CreateSystemField is the resolver for the createSystemField field.
func (r *mutationResolver) CreateSystemField(ctx context.Context, fieldCode string, fieldName string, dataType string, description *string, defaultValue *string) (*model.SystemField, error) {
	systemField := &model.SystemField{
		FieldCode:    fieldCode,
		FieldName:    fieldName,
		DataType:     dataType,
		Description:  description,
		DefaultValue: defaultValue,
	}

	trans := database.NewTransactorRepo(&gorm.DB{})
	repo := repositories.NewSystemFieldRepo(&gorm.DB{})
	systemFieldService := services.NewSystemFieldService(repo, trans)
	hld := gqlhandlers.NewGQLSystemFieldHandler(systemFieldService)
	d := hld.CreateSystemField(systemField)
	return d, nil
}

// UpdateSystemField is the resolver for the updateSystemField field.
func (r *mutationResolver) UpdateSystemField(ctx context.Context, id string, fieldCode *string, fieldName *string, dataType *string, description *string, defaultValue *string) (*model.SystemField, error) {
	panic(fmt.Errorf("not implemented: UpdateSystemField - updateSystemField"))
}

// DeleteSystemField is the resolver for the deleteSystemField field.
func (r *mutationResolver) DeleteSystemField(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteSystemField - deleteSystemField"))
}

// SystemField is the resolver for the systemField field.
func (r *queryResolver) SystemField(ctx context.Context, id string) (*model.SystemField, error) {
	panic(fmt.Errorf("not implemented: SystemField - systemField"))
}

// SystemFields is the resolver for the systemFields field.
func (r *queryResolver) SystemFields(ctx context.Context) ([]*model.SystemField, error) {
	panic(fmt.Errorf("not implemented: SystemFields - systemFields"))
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
