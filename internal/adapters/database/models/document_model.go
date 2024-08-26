package models

import (
	"time"

	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	ID           uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	DocumentName string         `json:"document_name"`
	Issuer       string         `json:"issuer"`
	OrderID      string         `json:"order_id"`
	Order        Order          `gorm:"foreignkey:OrderID"`
}

var TNDocument = "documents"

func (st *Document) TableName() string {
	return TNDocument
}

type DocumentTemplate struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	DocumentName    string `json:"document_name"`
	TemplateVersion string `json:"template_version"`
	IsDefault       bool   `json:"is_default"`
	IsEnable        bool   `json:"is_enable"`
}

var TNDocumentTemplate = "document_template"

func (st *DocumentTemplate) TableName() string {
	return TNDocumentTemplate
}

type DocumentVersion struct {
	gorm.Model
	ID         uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt  time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	DocumentID uint           `json:"document_id"`
	Document   Document       `gorm:"foreignkey:DocumentID"`

	DocumentTemplateID uint             `json:"document_template_id"`
	DocumentTemplate   DocumentTemplate `gorm:"foreignkey:DocumentTemplateID"`

	VersionNumber int `json:"version_number"`
}

var TNDocumentVersion = "document_versions"

func (st *DocumentVersion) TableName() string {
	return TNDocumentVersion
}

type DocumentTemplateField struct {
	gorm.Model
	ID                 uint             `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt          time.Time        `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time        `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt   `gorm:"index" json:"deleted_at,omitempty"`
	DocumentTemplateID uint             `json:"document_template_id"`
	DocumentTemplate   DocumentTemplate `gorm:"foreignkey:DocumentTemplateID"`
	SystemFieldID      uint             `json:"system_field_id"`
	SystemField        SystemField      `gorm:"foreignkey:SystemFieldID"`
}

var TNDocumentTemplateField = "document_template_fields"

func (st *DocumentTemplateField) TableName() string {
	return TNDocumentTemplateField
}

type DocumentVersionFieldValue struct {
	gorm.Model
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	DocumentVersionID uint            `json:"document_version_id"`
	DocumentVersion   DocumentVersion `gorm:"foreignkey:DocumentVersionID"`

	DocumentTemplateFieldID uint                  `json:"document_template_field_id"`
	DocumentTemplateField   DocumentTemplateField `gorm:"foreignkey:DocumentTemplateFieldID"`
	Value                   string                `json:"value"`
}

var TNDocumentVersionFieldValue = "document_version_field_value"

func (st *DocumentVersionFieldValue) TableName() string {
	return TNDocumentVersionFieldValue
}

type LogDocumentVersionFieldValue struct {
	gorm.Model
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	DocumentVersionFieldValueID uint                      `json:"document_version_field_value_id"`
	DocumentVersionFieldValue   DocumentVersionFieldValue `gorm:"foreignkey:DocumentVersionFieldValueID"`

	MasterFileID *uint       `json:"master_file_id"`
	MasterFile   *MasterFile `gorm:"foreignkey:MasterFileID"`

	PreviousValue *string `json:"previous_value"`
	ModifiedValue string  `json:"modified_value"`
}

var TNLogDocumentVersionFieldValue = "log_document_version_field_value"

func (st *LogDocumentVersionFieldValue) TableName() string {
	return TNLogDocumentVersionFieldValue
}
