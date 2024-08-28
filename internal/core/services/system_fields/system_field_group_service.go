package services

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database"
	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"

	domain "github.com/billowdev/exclusive-go-hexa/internal/core/domain/system_fields"
	ports "github.com/billowdev/exclusive-go-hexa/internal/core/ports/system_fields"
	"github.com/billowdev/exclusive-go-hexa/pkg/configs"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"github.com/billowdev/exclusive-go-hexa/pkg/utils"
)

type SystemFieldGroupServiceImpls struct {
	repo       ports.ISystemGroupFieldRepository
	transactor database.IDatabaseTransactor
}

func NewSystemGroupFieldService(
	repo ports.ISystemGroupFieldRepository,
	transactor database.IDatabaseTransactor,
) ports.ISystemGroupFieldService {
	return &SystemFieldGroupServiceImpls{repo: repo, transactor: transactor}
}

// CreateSystemGroupField implements ports.ISystemGroupFieldService.
func (s *SystemFieldGroupServiceImpls) CreateSystemGroupField(ctx context.Context, payload *models.SystemGroupField) utils.APIResponse {
	if err := s.repo.CreateSystemGroupField(ctx, payload); err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: nil}
}

// DeleteSystemGroupField implements ports.ISystemGroupFieldService.
func (s *SystemFieldGroupServiceImpls) DeleteSystemGroupField(ctx context.Context, id uint) utils.APIResponse {
	if err := s.repo.DeleteSystemGroupField(ctx, id); err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: nil}
}

// GetSystemGroupField implements ports.ISystemGroupFieldService.
func (s *SystemFieldGroupServiceImpls) GetSystemGroupField(ctx context.Context, id uint) utils.APIResponse {

	data, err := s.repo.GetSystemGroupField(ctx, id)
	if err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	if data == nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Not Found", Data: nil}
	}
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: data}
}

// GetSystemGroupFields implements ports.ISystemGroupFieldService.
func (s *SystemFieldGroupServiceImpls) GetSystemGroupFields(ctx context.Context) pagination.Pagination[[]domain.SystemGroupFieldDomain] {
	data, err := s.repo.GetSystemGroupFields(ctx)
	if err != nil {
		return pagination.Pagination[[]domain.SystemGroupFieldDomain]{}
	}
	// Convert repository data to domain models
	newData := utils.ConvertSlice(data.Rows, domain.ToSystemGroupFieldDomain)

	return pagination.Pagination[[]domain.SystemGroupFieldDomain]{
		Rows:       newData,
		Links:      data.Links,
		Total:      data.Total,
		Page:       data.Page,
		PageSize:   data.PageSize,
		TotalPages: data.TotalPages,
	}
}

// UpdateSystemGroupField implements ports.ISystemGroupFieldService.
func (s *SystemFieldGroupServiceImpls) UpdateSystemGroupField(ctx context.Context, payload *models.SystemGroupField) utils.APIResponse {
	err := s.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		err := s.repo.UpdateSystemGroupField(ctx, payload)
		if err != nil {
			return err // Transaction is rolled back if error occurred during deletion or update operation
		}
		return nil
	})
	if err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: nil}

}
