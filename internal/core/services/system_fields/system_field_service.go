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

type SystemFieldServiceImpls struct {
	repo       ports.ISystemFieldRepository
	transactor database.IDatabaseTransactor
}

func NewSystemFieldService(
	repo ports.ISystemFieldRepository,
	transactor database.IDatabaseTransactor,
) ports.ISystemFieldService {
	return &SystemFieldServiceImpls{
		repo:       repo,
		transactor: transactor,
	}
}

// CreateSystemField implements ports.ISystemFieldService.
func (s *SystemFieldServiceImpls) CreateSystemField(ctx context.Context, payload *models.SystemField) utils.APIResponse {
	// tx, err := s.transactor.BeginTransaction()
	// if err != nil {
	// 	return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	// }
	// ctxWithTx := database.InjectTx(ctx, tx) // Inject transaction into context
	if err := s.repo.CreateSystemField(ctx, payload); err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	res := domain.ToDomainModel(payload)
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: res}
}

// DeleteSystemField implements ports.ISystemFieldService.
func (s *SystemFieldServiceImpls) DeleteSystemField(ctx context.Context, id uint) utils.APIResponse {
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
func (s *SystemFieldServiceImpls) GetSystemField(ctx context.Context, id uint) utils.APIResponse {
	data, err := s.repo.GetSystemField(ctx, id)
	if err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	if data == nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: nil}
	}
	res := domain.ToDomainModel(data)
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: res}
}

// GetSystemFields implements ports.ISystemFieldService.
func (s *SystemFieldServiceImpls) GetSystemFields(ctx context.Context) pagination.Pagination[[]domain.SystemFieldDomain] {
	data, err := s.repo.GetSystemFields(ctx)
	if err != nil {
		return pagination.Pagination[[]domain.SystemFieldDomain]{}
	}
	// Convert repository data to domain models
	newData := utils.ConvertSlice(data.Rows, domain.ToDomainModel)

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
func (s *SystemFieldServiceImpls) UpdateSystemField(ctx context.Context, payload *models.SystemField) utils.APIResponse {
	// tx, err := s.transactor.BeginTransaction()
	// if err!= nil {
	//     return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	// }
	// ctxWithTx := database.InjectTx(ctx, tx) // Inject transaction into context
	if err := s.repo.UpdateSystemField(ctx, payload); err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	res := domain.ToDomainModel(payload)
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: res}
}
