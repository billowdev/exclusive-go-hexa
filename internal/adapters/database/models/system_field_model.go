package models

import (
	"time"

	"gorm.io/gorm"
)

type SystemField struct {
	gorm.Model
	ID           uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	FieldCode    string         `json:"field_code"`
	FieldName    string         `json:"field_name"`
	DataType     string         `json:"data_type"`
	Description  string         `json:"description"`
	DefaultValue string         `json:"default_value"`
}

var TNSystemField = "system_fields"

func (st *SystemField) TableName() string {
	return TNSystemField
}
