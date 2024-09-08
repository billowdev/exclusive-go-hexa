package domain

import (
	"time"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
)

type DocumentDomain struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DocumentName string    `json:"document_name"`
	Issuer       string    `json:"issuer"`
	OrderID      string    `json:"order_id"`
}

func ToDocumentDomain(data *models.Document) DocumentDomain {
	if data == nil {
		return DocumentDomain{}
	}
	return DocumentDomain{
		ID:           data.ID,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
		DocumentName: data.DocumentName,
		Issuer:       data.Issuer,
		OrderID:      data.OrderID,
	}
}

func FromToDocumentModel(data DocumentDomain) *models.Document {
	return &models.Document{
		ID:           data.ID,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
		DocumentName: data.DocumentName,
		Issuer:       data.Issuer,
		OrderID:      data.OrderID,
	}
}
