package repositories

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	ports "github.com/billowdev/exclusive-go-hexa/internal/core/ports/activities"
	"gorm.io/gorm"
)

type ActivityRepositoryImpls struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) ports.IActivityRepository {
	return &ActivityRepositoryImpls{db: db}
}

func (r *ActivityRepositoryImpls) CreateActivity(ctx context.Context, activity *models.Activity) error {
	return r.db.WithContext(ctx).Create(activity).Error
}
