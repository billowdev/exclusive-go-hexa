package repositories

import (
	"context"

	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/pagination"
	"github.com/stretchr/testify/mock"
)

type MockSystemFieldRepository struct {
	mock.Mock
}

func (m *MockSystemFieldRepository) CreateSystemField(ctx context.Context, payload *models.SystemField) error {
	args := m.Called(ctx, payload)
	return args.Error(0)
}

func (m *MockSystemFieldRepository) DeleteSystemField(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockSystemFieldRepository) GetSystemField(ctx context.Context, id uint) (*models.SystemField, error) {
	args := m.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.SystemField), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockSystemFieldRepository) GetSystemFields(ctx context.Context) (*pagination.Pagination[[]models.SystemField], error) {
	args := m.Called(ctx)
	if args.Get(0) != nil {
		return args.Get(0).(*pagination.Pagination[[]models.SystemField]), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockSystemFieldRepository) UpdateSystemField(ctx context.Context, payload *models.SystemField) error {
	args := m.Called(ctx, payload)
	return args.Error(0)
}
