package repositories

import (
	"context"
	"testing"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateSystemField(t *testing.T) {
	mockRepo := new(MockSystemFieldRepository)
	payload := &models.SystemField{
		ID:           1,
		FieldCode:    "FC001",
		FieldName:    "Test Field",
		DataType:     "string",
		Description:  "A test field",
		DefaultValue: "default",
	}

	// Set up expectations
	mockRepo.On("CreateSystemField", mock.Anything, payload).Return(nil)

	// Call the method
	err := mockRepo.CreateSystemField(context.Background(), payload)

	// Assert the results
	assert.NoError(t, err)

	// Verify expectations
	mockRepo.AssertExpectations(t)
}

func TestDeleteSystemField(t *testing.T) {
	mockRepo := new(MockSystemFieldRepository)
	id := uint(1)

	// Set up expectations
	mockRepo.On("DeleteSystemField", mock.Anything, id).Return(nil)

	// Call the method
	err := mockRepo.DeleteSystemField(context.Background(), id)

	// Assert the results
	assert.NoError(t, err)

	// Verify expectations
	mockRepo.AssertExpectations(t)
}

func TestGetSystemField(t *testing.T) {
	mockRepo := new(MockSystemFieldRepository)
	id := uint(1)
	expectedField := &models.SystemField{
		ID:           id,
		FieldCode:    "FC001",
		FieldName:    "Test Field",
		DataType:     "string",
		Description:  "A test field",
		DefaultValue: "default",
	}

	// Set up expectations
	mockRepo.On("GetSystemField", mock.Anything, id).Return(expectedField, nil)

	// Call the method
	field, err := mockRepo.GetSystemField(context.Background(), id)

	// Assert the results
	assert.NoError(t, err)
	assert.Equal(t, expectedField, field)

	// Verify expectations
	mockRepo.AssertExpectations(t)
}

func TestGetSystemFields(t *testing.T) {
	mockRepo := new(MockSystemFieldRepository)
	expectedPagination := pagination.Pagination[[]models.SystemField]{
		Links: pagination.PaginationLinks{
			Next:     "",
			Previous: "",
		},
		Total:      2,
		Page:       1,
		PageSize:   10,
		TotalPages: 1,
		Rows: []models.SystemField{
			{ID: 1, FieldCode: "FC001", FieldName: "Field 1"},
			{ID: 2, FieldCode: "FC002", FieldName: "Field 2"},
		},
	}

	// Set up expectations
	mockRepo.On("GetSystemFields", mock.Anything).Return(&expectedPagination, nil)

	// Call the method
	pagination, err := mockRepo.GetSystemFields(context.Background())

	// Assert the results
	assert.NoError(t, err)
	assert.Equal(t, expectedPagination, *pagination)

	// Verify expectations
	mockRepo.AssertExpectations(t)
}

func TestUpdateSystemField(t *testing.T) {
	mockRepo := new(MockSystemFieldRepository)
	payload := &models.SystemField{
		ID:           1,
		FieldCode:    "FC001",
		FieldName:    "Updated Field",
		DataType:     "string",
		Description:  "Updated description",
		DefaultValue: "new default",
	}

	// Set up expectations
	mockRepo.On("UpdateSystemField", mock.Anything, payload).Return(nil)

	// Call the method
	err := mockRepo.UpdateSystemField(context.Background(), payload)

	// Assert the results
	assert.NoError(t, err)

	// Verify expectations
	mockRepo.AssertExpectations(t)
}
