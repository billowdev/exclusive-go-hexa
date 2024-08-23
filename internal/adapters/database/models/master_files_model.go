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

type SystemGroupField struct {
	gorm.Model
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	GroupName   string         `json:"group_name"`
	Description string         `json:"description"`
}

var TNSystemGroupField = "system_group_fields"

func (st *SystemGroupField) TableName() string {
	return TNSystemGroupField
}

type ConfigSystemMasterFileField struct {
	ID                 uint             `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt          time.Time        `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time        `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt   `gorm:"index" json:"deleted_at,omitempty"`
	SystemFieldID      string           `json:"system_field_id"`
	SystemField        SystemField      `gorm:"foreignkey:SystemFieldID"`
	SystemGroupFieldID string           `json:"system_group_field_id"`
	SystemGroupField   SystemGroupField `gorm:"foreignkey:SystemGroupFieldID"`
}

type MasterFile struct {
	gorm.Model
	ID            uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt     time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	SystemFieldID uint           `json:"system_field_id"`
	SystemField   SystemField    `gorm:"foreignkey:SystemFieldID"`
	Value         string         `json:"value"`
}

var TNMasterFile = "master_files"

func (st *MasterFile) TableName() string {
	return TNMasterFile
}

type LogMasterFile struct {
	ID            uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt     time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	MasterFileID  uint           `json:"master_file_id"`
	MasterFile    MasterFile     `gorm:"foreignkey:MasterFileID"`
	PreviousValue string         `json:"previous_value"`
	ModifiedValue string         `json:"modified_value"`
}

var TNLogMasterFile = "log_master_files"

func (st *LogMasterFile) TableName() string {
	return TNLogMasterFile
}
