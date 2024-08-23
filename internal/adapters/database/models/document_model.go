package models

import (
	"time"

	"gorm.io/gorm"
)

type DocumentModel struct {
	gorm.Model
	ID           string     `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	DocumentName string     `json:"document_name"`
	Issuer       string     `json:"issuer"`
	OrderID      string     `json:"order_id"`
	Order        OrderModel `gorm:"foreignkey:OrderID"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

var TNDocument = "documents"

func (st *DocumentModel) TableName() string {
	return TNDocument
}

type DocumentVersionModel struct {
	gorm.Model
	ID            string        `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	DocumentID    string        `json:"document_id"`
	Document      DocumentModel `gorm:"foreignkey:DocumentID"`
	VersionNumber int           `json:"version_number"`
	CreatedAt     time.Time     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
}

var TNDocumentVersion = "document_versions"

func (st *DocumentVersionModel) TableName() string {
	return TNDocumentVersion
}
