package seeders

import (
	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	"gorm.io/gorm"
)

var SEED_DOCUMENT = []byte(`[
	{
		"id": 1,
		"document_name": "BILL OF LADING",
		"issuer":"LINER",
		"order_id": "1"
	},
	{
		"id": 2,
		"document_name": "INVOICE",
        "issuer":"CUSTOMER",
        "order_id": "2"
    },
	{
		"id":3,
        "document_name": "ORDER FORM",
        "issuer":"SUPPLIER",
        "order_id": "3"
    }
]`)

func SeedDocument(db *gorm.DB) error {
	var data []models.Document
	seed := SEED_DOCUMENT
	if err := BaseSeeder(db, &data, seed); err != nil {
		return err
	}
	return nil
}

var SEED_DOCUMENT_TEMPLATE = []byte(`[
	{
		"document_name": "BILL OF LADING",
		"template_version": "1.0",
		"is_default": true,
        "is_enable": true
	},
	{
		"document_name": "INVOICE",
		"template_version": "1.0",
		"is_default": true,
        "is_enable": true
	}, 
	{
		"document_name": "ORDER FORM",
		"template_version": "1.0",
		"is_default": true,
        "is_enable": true
	}
]`)

func SeedDocumentTemplate(db *gorm.DB) error {
	var data []models.DocumentTemplate
	seed := SEED_DOCUMENT_TEMPLATE
	if err := BaseSeeder(db, &data, seed); err != nil {
		return err
	}
	return nil
}

var SEED_DOCUMENT_TEMPLATE_FIELD = []byte(`[
	{
		"id":1,
		"document_template_id": 1,
		"system_field_id": 1
	},
	{
		"id":2,
		"document_template_id": 1,
		"system_field_id": 2
	},
	{
		"id":3,
		"document_template_id": 1,
		"system_field_id": 3
	},
	{
		"id":4,
		"document_template_id": 2,
		"system_field_id": 1
	},
	{
		"id": 5,
		"document_template_id": 2,
		"system_field_id": 2
	},
	{
		"id":6,
		"document_template_id": 2,
		"system_field_id": 3
	}
]`)

func SeedDocumentTemplateField(db *gorm.DB) error {
	var data []models.DocumentTemplateField
	seed := SEED_DOCUMENT_TEMPLATE_FIELD
	if err := BaseSeeder(db, &data, seed); err != nil {
		return err
	}
	return nil
}
