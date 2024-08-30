package services

import (
	"context"
	"time"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database"
	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	domain "github.com/billowdev/exclusive-go-hexa/internal/core/domain/system_fields"
	ports "github.com/billowdev/exclusive-go-hexa/internal/core/ports/system_fields"
	"github.com/billowdev/exclusive-go-hexa/pkg/configs"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"github.com/billowdev/exclusive-go-hexa/pkg/utils"
)

type SystemFieldServiceImpl struct {
	repo       ports.ISystemFieldRepository
	transactor database.IDatabaseTransactor
}

func NewSystemFieldService(
	repo ports.ISystemFieldRepository,
	transactor database.IDatabaseTransactor,
) ports.ISystemFieldService {
	return &SystemFieldServiceImpl{
		repo:       repo,
		transactor: transactor,
	}
}

// CreateSystemField implements ports.ISystemFieldService.
func (s *SystemFieldServiceImpl) CreateSystemField(ctx context.Context, payload *models.SystemField) utils.APIResponse {
	var result utils.APIResponse

	// Use the WithTransactionContextTimeout function to handle the transaction
	err := s.transactor.WithTransactionContextTimeout(ctx, 5*time.Second, func(txCtx context.Context) error {
		// Create the system field entry in the database using the transaction context
		if err := s.repo.CreateSystemField(txCtx, payload); err != nil {
			return err // Return error to trigger rollback
		}

		// If no errors occurred, prepare a successful response
		result = utils.APIResponse{
			StatusCode:    configs.API_SUCCESS_CODE,
			StatusMessage: "Success",
			Data:          domain.ToSystemFieldDomain(payload),
		}
		return nil // Indicate success
	})

	// Check if there was an error during the transaction
	if err != nil {
		// Prepare an error response if something went wrong
		result = utils.APIResponse{
			StatusCode:    configs.API_ERROR_CODE,
			StatusMessage: "Error",
			Data:          err,
		}
	}

	// Return the result of the transaction
	return result
}

// DeleteSystemField implements ports.ISystemFieldService.
func (s *SystemFieldServiceImpl) DeleteSystemField(ctx context.Context, id uint) utils.APIResponse {
	err := s.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		err := s.repo.DeleteSystemField(txCtx, id)
		if err != nil {
			return err
		}
		return nil // Transaction is successful if no error occurred during deletion
	})

	if err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}

	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: nil}
}

// GetSystemField implements ports.ISystemFieldService.
func (s *SystemFieldServiceImpl) GetSystemField(ctx context.Context, id uint) utils.APIResponse {
	data, err := s.repo.GetSystemField(ctx, id)
	if err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	if data == nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: nil}
	}
	res := domain.ToSystemFieldDomain(data)
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: res}
}

// GetSystemFields implements ports.ISystemFieldService.
func (s *SystemFieldServiceImpl) GetSystemFields(ctx context.Context) pagination.Pagination[[]domain.SystemFieldDomain] {
	data, err := s.repo.GetSystemFields(ctx)
	if err != nil {
		return pagination.Pagination[[]domain.SystemFieldDomain]{}
	}
	// Convert repository data to domain models
	newData := utils.ConvertSlice(data.Rows, domain.ToSystemFieldDomain)

	return pagination.Pagination[[]domain.SystemFieldDomain]{
		Rows:       newData,
		Links:      data.Links,
		Total:      data.Total,
		Page:       data.Page,
		PageSize:   data.PageSize,
		TotalPages: data.TotalPages,
	}
}

// UpdateSystemField implements ports.ISystemFieldService.
func (s *SystemFieldServiceImpl) UpdateSystemField(ctx context.Context, payload *models.SystemField) utils.APIResponse {
	var result utils.APIResponse

	// Use the WithinTransaction function to handle the transaction
	err := s.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		// Update the system field entry in the database using the transaction context
		if err := s.repo.UpdateSystemField(txCtx, payload); err != nil {
			return err // Return error to trigger rollback
		}

		// If no errors occurred, prepare a successful response
		result = utils.APIResponse{
			StatusCode:    configs.API_SUCCESS_CODE,
			StatusMessage: "Success",
			Data:          domain.ToSystemFieldDomain(payload),
		}
		return nil // Indicate success
	})

	// Check if there was an error during the transaction
	if err != nil {
		// Prepare an error response if something went wrong
		result = utils.APIResponse{
			StatusCode:    configs.API_ERROR_CODE,
			StatusMessage: "Error",
			Data:          err,
		}
	}

	// Return the result of the transaction
	return result
}
