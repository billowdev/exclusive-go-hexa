package domain

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database"
	domain "github.com/billowdev/exclusive-go-hexa/internal/core/domain/documents"
	ports "github.com/billowdev/exclusive-go-hexa/internal/core/ports/documents"
	"github.com/billowdev/exclusive-go-hexa/pkg/configs"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"github.com/billowdev/exclusive-go-hexa/pkg/utils"
)

type DocumentServiceImpl struct {
	transactor database.IDatabaseTransactor
	repo       ports.IDocumentRepository
}

func NewDocumentService(transactor database.IDatabaseTransactor, repo ports.IDocumentRepository) ports.IDocumentService {
	return &DocumentServiceImpl{transactor, repo}
}

// CreateDocument implements ports.IDocumentService.
func (d *DocumentServiceImpl) CreateDocument(ctx context.Context, payload domain.DocumentDomain) utils.APIResponse {
	data := domain.FromToDocumentModel(payload)
	if err := d.repo.CreateDocument(ctx, data); err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	res := domain.ToDocumentDomain(data)
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: res}
}

// DeleteDocument implements ports.IDocumentService.
func (d *DocumentServiceImpl) DeleteDocument(ctx context.Context, id uint) utils.APIResponse {
	if err := d.repo.DeleteDocument(ctx, id); err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: nil}
}

// GetDocument implements ports.IDocumentService.
func (d *DocumentServiceImpl) GetDocument(ctx context.Context, id uint) utils.APIResponse {
	data, err := d.repo.GetDocument(ctx, id)
	if err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	if data == nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Not Found", Data: nil}
	}
	res := domain.ToDocumentDomain(data)
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: res}
}

// GetDocuments implements ports.IDocumentService.
func (d *DocumentServiceImpl) GetDocuments(ctx context.Context) pagination.PaginationResponse {
	data, err := d.repo.GetDocuments(ctx)
	if err != nil {
		return pagination.PaginationResponse{}
	}
	// Convert repository data to domain models
	newData := utils.ConvertSlice(data.Rows, domain.ToDocumentDomain)
	return pagination.PaginationResponse{
		StatusCode:    configs.API_SUCCESS_CODE,
		StatusMessage: "Success",
		Data:          newData,
		Pagination: pagination.PaginationInfo{
			Links:      data.Links,
			Total:      data.Total,
			Page:       data.Page,
			PageSize:   data.PageSize,
			TotalPages: data.TotalPages,
		},
	}
}

// UpdateDocument implements ports.IDocumentService.
func (d *DocumentServiceImpl) UpdateDocument(ctx context.Context, payload domain.DocumentDomain) utils.APIResponse {
	data := domain.FromToDocumentModel(payload)
	if err := d.repo.UpdateDocument(ctx, data); err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	res := domain.ToDocumentDomain(data)
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: res}
}
