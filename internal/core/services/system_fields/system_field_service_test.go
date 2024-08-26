package services

import (
	"context"
	"testing"

	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	"github.com/billowdev/document-system-field-manager/pkg/configs"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/pagination"
	"github.com/billowdev/document-system-field-manager/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateSystemField(t *testing.T) {
	mockService := new(MockSystemFieldService)
	payload := &models.SystemField{
		ID:           1,
		FieldCode:    "FC001",
		FieldName:    "Test Field",
		DataType:     "string",
		Description:  "A test field",
		DefaultValue: "default",
	}

	// Set up expectations
	mockService.On("CreateSystemField", mock.Anything, payload).Return(utils.APIResponse{
		StatusCode:    "success",
		StatusMessage: "Success",
		Data:          payload,
	})

	// Call the method
	response := mockService.CreateSystemField(context.Background(), payload)

	// Assert the results
	assert.Equal(t, configs.API_SUCCESS_CODE, response.StatusCode)
	assert.Equal(t, "Success", response.StatusMessage)
	assert.Equal(t, payload, response.Data)

	// Verify expectations
	mockService.AssertExpectations(t)
}

func TestDeleteSystemField(t *testing.T) {
	mockService := new(MockSystemFieldService)
	id := uint(1)

	// Set up expectations
	mockService.On("DeleteSystemField", mock.Anything, id).Return(utils.APIResponse{
		StatusCode:    "success",
		StatusMessage: "Success",
		Data:          nil,
	})

	// Call the method
	response := mockService.DeleteSystemField(context.Background(), id)

	// Assert the results
	assert.Equal(t, configs.API_SUCCESS_CODE, response.StatusCode)
	assert.Equal(t, "Success", response.StatusMessage)
	assert.Nil(t, response.Data)

	// Verify expectations
	mockService.AssertExpectations(t)
}

func TestGetSystemField(t *testing.T) {
	mockService := new(MockSystemFieldService)
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
	mockService.On("GetSystemField", mock.Anything, id).Return(utils.APIResponse{
		StatusCode:    "success",
		StatusMessage: "Success",
		Data:          expectedField,
	})

	// Call the method
	response := mockService.GetSystemField(context.Background(), id)

	// Assert the results
	assert.Equal(t, configs.API_SUCCESS_CODE, response.StatusCode)
	assert.Equal(t, "Success", response.StatusMessage)
	assert.Equal(t, expectedField, response.Data)

	// Verify expectations
	mockService.AssertExpectations(t)
}

func TestGetSystemFields(t *testing.T) {
	mockService := new(MockSystemFieldService)
	expectedPagination := pagination.Pagination[[]models.SystemField]{
		Rows: []models.SystemField{
			{ID: 1, FieldCode: "FC001", FieldName: "Field 1"},
			{ID: 2, FieldCode: "FC002", FieldName: "Field 2"},
		},
		Links: pagination.PaginationLinks{
			Next:     "",
			Previous: "",
		},
		Total:      2,
		Page:       1,
		PageSize:   10,
		TotalPages: 1,
	}

	// Set up expectations
	mockService.On("GetSystemFields", mock.Anything).Return(expectedPagination)

	// Call the method
	response := mockService.GetSystemFields(context.Background())

	// Assert the results
	assert.Equal(t, expectedPagination.Rows, response.Rows)
	assert.Equal(t, expectedPagination.Links, response.Links)
	assert.Equal(t, expectedPagination.Total, response.Total)
	assert.Equal(t, expectedPagination.Page, response.Page)
	assert.Equal(t, expectedPagination.PageSize, response.PageSize)
	assert.Equal(t, expectedPagination.TotalPages, response.TotalPages)

	// Verify expectations
	mockService.AssertExpectations(t)
}

func TestUpdateSystemField(t *testing.T) {
	mockService := new(MockSystemFieldService)
	payload := &models.SystemField{
		ID:           1,
		FieldCode:    "FC001",
		FieldName:    "Updated Field",
		DataType:     "string",
		Description:  "Updated description",
		DefaultValue: "new default",
	}

	// Set up expectations
	mockService.On("UpdateSystemField", mock.Anything, payload).Return(utils.APIResponse{
		StatusCode:    "success",
		StatusMessage: "Success",
		Data:          payload,
	})

	// Call the method
	response := mockService.UpdateSystemField(context.Background(), payload)

	// Assert the results
	assert.Equal(t, configs.API_SUCCESS_CODE, response.StatusCode)
	assert.Equal(t, "Success", response.StatusMessage)
	assert.Equal(t, payload, response.Data)

	// Verify expectations
	mockService.AssertExpectations(t)
}
