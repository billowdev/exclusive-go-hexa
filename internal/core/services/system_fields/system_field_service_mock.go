package services

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"github.com/billowdev/exclusive-go-hexa/pkg/utils"
	"github.com/stretchr/testify/mock"
)

type MockSystemFieldService struct {
	mock.Mock
}

func (m *MockSystemFieldService) CreateSystemField(ctx context.Context, payload *models.SystemField) utils.APIResponse {
	args := m.Called(ctx, payload)
	return args.Get(0).(utils.APIResponse)
}

func (m *MockSystemFieldService) DeleteSystemField(ctx context.Context, id uint) utils.APIResponse {
	args := m.Called(ctx, id)
	return args.Get(0).(utils.APIResponse)
}

func (m *MockSystemFieldService) GetSystemField(ctx context.Context, id uint) utils.APIResponse {
	args := m.Called(ctx, id)
	return args.Get(0).(utils.APIResponse)
}

func (m *MockSystemFieldService) GetSystemFields(ctx context.Context) pagination.Pagination[[]models.SystemField] {
	args := m.Called(ctx)
	return args.Get(0).(pagination.Pagination[[]models.SystemField])
}

func (m *MockSystemFieldService) UpdateSystemField(ctx context.Context, payload *models.SystemField) utils.APIResponse {
	args := m.Called(ctx, payload)
	return args.Get(0).(utils.APIResponse)
}
