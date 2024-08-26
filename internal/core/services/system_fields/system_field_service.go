package services

import (
	"context"

	"github.com/billowdev/document-system-field-manager/internal/adapters/database"
	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	ports "github.com/billowdev/document-system-field-manager/internal/core/ports/system_fields"
	"github.com/billowdev/document-system-field-manager/pkg/configs"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/pagination"
	"github.com/billowdev/document-system-field-manager/pkg/utils"
)

type SystemFieldServiceImpls struct {
	repo       ports.ISystemFieldRepository
	transactor database.ITransactor
}

func NewSystemFieldService(
	repo ports.ISystemFieldRepository,
	transactor database.ITransactor,
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
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: payload}
}

// DeleteSystemField implements ports.ISystemFieldService.
func (s *SystemFieldServiceImpls) DeleteSystemField(ctx context.Context, id uint) utils.APIResponse {
	err := s.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		panic("test")
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
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: data}
}

// GetSystemFields implements ports.ISystemFieldService.
func (s *SystemFieldServiceImpls) GetSystemFields(ctx context.Context) pagination.Pagination[[]models.SystemField] {
	data, err := s.repo.GetSystemFields(ctx)
	if err != nil {
		return pagination.Pagination[[]models.SystemField]{}
	}
	return pagination.Pagination[[]models.SystemField]{
		Rows:       data.Rows,
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
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: payload}
}
