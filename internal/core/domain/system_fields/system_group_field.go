package domain

import (
	"time"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
)

type SystemGroupFieldDomain struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func ToSystemGroupFieldDomain(sf *models.SystemGroupField) SystemGroupFieldDomain {
	if sf == nil {
		return SystemGroupFieldDomain{}
	}

	return SystemGroupFieldDomain{
		ID:          sf.ID,
		CreatedAt:   sf.CreatedAt,
		UpdatedAt:   sf.UpdatedAt,
		Name:        sf.Name,
		Description: sf.Description,
	}
}
