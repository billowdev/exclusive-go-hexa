package repositories

import (
	"context"
	"testing"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/filters"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSystemFieldGroupRepository struct {
	mock.Mock
}

// CreateSystemGroupField implements ports.ISystemGroupFieldRepository.
func (m *MockSystemFieldGroupRepository) CreateSystemGroupField(ctx context.Context, payload *models.SystemGroupField) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// DeleteSystemGroupField implements ports.ISystemGroupFieldRepository.
func (m *MockSystemFieldGroupRepository) DeleteSystemGroupField(ctx context.Context, id uint) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// GetSystemGroupField implements ports.ISystemGroupFieldRepository.
func (m *MockSystemFieldGroupRepository) GetSystemGroupField(ctx context.Context, id uint) (*models.SystemGroupField, error) {
	args := m.Called(ctx)
	if args.Get(0) != nil {
		return args.Get(0).(*models.SystemGroupField), args.Error(1)
	}
	return nil, args.Error(1)
}

// GetSystemGroupFields implements ports.ISystemGroupFieldRepository.
func (m *MockSystemFieldGroupRepository) GetSystemGroupFields(ctx context.Context, p pagination.PaginationParams[filters.SystemGroupFieldFilter]) (*pagination.Pagination[[]models.SystemGroupField], error) {
	args := m.Called(ctx)
	if args.Get(0) != nil {
		return args.Get(0).(*pagination.Pagination[[]models.SystemGroupField]), args.Error(1)
	}
	return nil, args.Error(1)
}

// UpdateSystemGroupField implements ports.ISystemGroupFieldRepository.
func (m *MockSystemFieldGroupRepository) UpdateSystemGroupField(ctx context.Context, payload *models.SystemGroupField) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func TestCreateSystemGroupField(t *testing.T) {
	mockRepo := new(MockSystemFieldGroupRepository)
	payload := &models.SystemGroupField{
		ID:          1,
		Name:        "Group1",
		Description: "Test Group Field",
	}
	mockRepo.On("CreateSystemGroupField", context.Background(), payload).Return(nil)
	err := mockRepo.CreateSystemGroupField(context.Background(), payload)
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteSystemGroupField(t *testing.T) {
	mockRepo := new(MockSystemFieldGroupRepository)
	id := uint(1)
	mockRepo.On("DeleteSystemGroupField", context.Background(), id).Return(nil)
	err := mockRepo.DeleteSystemGroupField(context.Background(), id)
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateSystemGroupField(t *testing.T) {
	mockRepo := new(MockSystemFieldGroupRepository)
	payload := &models.SystemGroupField{
		ID:          1,
		Name:        "Updated Group",
		Description: "Updated Test Group Field",
	}
	mockRepo.On("UpdateSystemGroupField", context.Background(), payload).Return(nil)
	err := mockRepo.UpdateSystemGroupField(context.Background(), payload)
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetSystemGroupField(t *testing.T) {
	mockRepo := new(MockSystemFieldGroupRepository)
	id := uint(1)
	expectedField := &models.SystemGroupField{
		ID:          1,
		Name:        "Group1",
		Description: "Test Group Field",
	}
	mockRepo.On("GetSystemGroupField", context.Background(), id).Return(expectedField, nil)
	field, err := mockRepo.GetSystemGroupField(context.Background(), id)
	assert.Nil(t, err)
	assert.Equal(t, expectedField, field)
	mockRepo.AssertExpectations(t)
}

func TestGetSystemGroupFields(t *testing.T) {
	mockRepo := new(MockSystemFieldGroupRepository)
	paginationParams := pagination.PaginationParams[filters.SystemGroupFieldFilter]{
		Page:  1,
		Limit: 10,
		Filters: filters.SystemGroupFieldFilter{
			Name: "Test",
		},
	}
	expectedFields := []models.SystemGroupField{
		{
			ID:          1,
			Name:        "Group1",
			Description: "Test Group Field",
		},
	}
	ctx := context.Background()
	// pagination.SetFilters[]()
	mockRepo.On("GetSystemGroupFields", ctx, paginationParams).Return(&pagination.Pagination[[]models.SystemGroupField]{
		Rows:       expectedFields,
		TotalPages: 1,
		Total:      1,
	}, nil)
	fields, err := mockRepo.GetSystemGroupFields(ctx, paginationParams)
	assert.Nil(t, err)
	assert.Equal(t, expectedFields, fields.Rows)
}
