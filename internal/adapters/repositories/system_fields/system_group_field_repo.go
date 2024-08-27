package repositories

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	ports "github.com/billowdev/exclusive-go-hexa/internal/core/ports/system_fields"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/filters"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"gorm.io/gorm"
)

type SystemGroupFieldRepositoryImpls struct {
	db *gorm.DB
}

func NewSystemGroupFieldRepository(db *gorm.DB) ports.ISystemGroupFieldRepository {
	return &SystemGroupFieldRepositoryImpls{db: db}
}

// CreateSystemGroupField implements ports.ISystemGroupFieldRepository.
func (s *SystemGroupFieldRepositoryImpls) CreateSystemGroupField(ctx context.Context, payload *models.SystemGroupField) error {
	panic("unimplemented")
}

// DeleteSystemGroupField implements ports.ISystemGroupFieldRepository.
func (s *SystemGroupFieldRepositoryImpls) DeleteSystemGroupField(ctx context.Context, id uint) error {
	panic("unimplemented")
}

// GetSystemGroupField implements ports.ISystemGroupFieldRepository.
func (s *SystemGroupFieldRepositoryImpls) GetSystemGroupField(ctx context.Context, id uint) (*models.SystemGroupField, error) {
	panic("unimplemented")
}

// GetSystemGroupFields implements ports.ISystemGroupFieldRepository.
func (s *SystemGroupFieldRepositoryImpls) GetSystemGroupFields(ctx context.Context, p pagination.PaginationParams[filters.SystemGroupFieldFilter]) (*pagination.Pagination[[]models.SystemGroupField], error) {
	panic("unimplemented")
}

// UpdateSystemGroupField implements ports.ISystemGroupFieldRepository.
func (s *SystemGroupFieldRepositoryImpls) UpdateSystemGroupField(ctx context.Context, payload *models.SystemGroupField) error {
	panic("unimplemented")
}
