package domain

import (
	"time"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	"gorm.io/gorm"
)

type SystemFieldDomain struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    *time.Time
	FieldCode    string `json:"field_code"`
	FieldName    string `json:"field_name"`
	DataType     string `json:"data_type"`
	Description  string `json:"description"`
	DefaultValue string `json:"default_value"`
}

// func ToDomainModel(sf *models.SystemField) *SystemFieldDomain {
// 	if sf == nil {
// 		return nil
// 	}

//		var deletedAt *time.Time
//		if sf.DeletedAt.Valid {
//			t := sf.DeletedAt.Time
//			deletedAt = &t
//		}
//		return &SystemFieldDomain{
//			ID:           sf.ID,
//			CreatedAt:    sf.CreatedAt,
//			UpdatedAt:    sf.UpdatedAt,
//			DeletedAt:    deletedAt,
//			FieldCode:    sf.FieldCode,
//			FieldName:    sf.FieldName,
//			DataType:     sf.DataType,
//			Description:  sf.Description,
//			DefaultValue: sf.DefaultValue,
//		}
//	}
//
// Adjusted ToDomainModel function to accept a pointer
func ToSystemFieldDomain(sf *models.SystemField) SystemFieldDomain {
	if sf == nil {
		return SystemFieldDomain{}
	}

	deletedAt := sf.DeletedAt.Time
	return SystemFieldDomain{
		ID:           sf.ID,
		CreatedAt:    sf.CreatedAt,
		UpdatedAt:    sf.UpdatedAt,
		DeletedAt:    &deletedAt,
		FieldCode:    sf.FieldCode,
		FieldName:    sf.FieldName,
		DataType:     sf.DataType,
		Description:  sf.Description,
		DefaultValue: sf.DefaultValue,
	}
}

func FromDomainModel(sfd *SystemFieldDomain) *models.SystemField {
	if sfd == nil {
		return nil
	}

	var deletedAt gorm.DeletedAt
	if sfd.DeletedAt != nil {
		deletedAt = gorm.DeletedAt{
			Time:  *sfd.DeletedAt,
			Valid: true,
		}
	} else {
		deletedAt = gorm.DeletedAt{
			Valid: false,
		}
	}

	return &models.SystemField{
		ID:           sfd.ID,
		CreatedAt:    sfd.CreatedAt,
		UpdatedAt:    sfd.UpdatedAt,
		DeletedAt:    deletedAt,
		FieldCode:    sfd.FieldCode,
		FieldName:    sfd.FieldName,
		DataType:     sfd.DataType,
		Description:  sfd.Description,
		DefaultValue: sfd.DefaultValue,
	}
}
