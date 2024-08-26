package seeders

import (
	"fmt"

	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	"gorm.io/gorm"
)

var SEED_DOCUMENT_VERSION = []byte(`[
    {
        "document_id": 1,
        "document_template_id": 1,
        "version_number": 1
    },
    {
        "document_id": 2,
        "document_template_id": 2,
        "version_number": 1
    }
]`)

func SeedDocumentVersion(db *gorm.DB) error {
	var data []models.DocumentVersion
	seed := SEED_DOCUMENT_VERSION
	if err := BaseSeeder(db, &data, seed); err != nil {
		fmt.Println("---SeedDocumentVersion---")
		fmt.Println(err)
		fmt.Println("------")
		return err
	}
	return nil
}

var SEED_DOCUMENT_VERSION_FIELD_VALUE = []byte(`[
	{
		"id": 1,
		"document_version_id": 1,
		"document_template_field_id": 1,
		"value": "BN010011"
	},
	{
		"id": 2,
		"document_version_id": 1,
		"document_template_field_id": 2,
		"value": "THLCH"
	},
	{
		"id": 3,
		"document_version_id": 1,
		"document_template_field_id": 3,
		"value": "SGSIN"
	},
	{
		"id": 4,
		"document_version_id": 2,
		"document_template_field_id": 4,
		"value": "BN020022"
	},
	{
		"id": 5,
		"document_version_id": 2,
		"document_template_field_id": 5,
		"value": "SGSIN"
	},
	{
		"id": 6,
		"document_version_id": 2,
		"document_template_field_id": 6,
		"value": "THLCH"
	}
]`)

func SeedDocumentVersionFieldValue(db *gorm.DB) error {
	var data []models.DocumentVersionFieldValue
	seed := SEED_DOCUMENT_VERSION_FIELD_VALUE
	if err := BaseSeeder(db, &data, seed); err != nil {
		fmt.Println("---SeedDocumentVersionFieldValue---")
		fmt.Println(err)
		fmt.Println("------")
		return err
	}
	return nil
}

var SEED_LOG_DOCUMENT_VERSION_FIELD_VALUE = []byte(`[
	{
		"id": 1,
		"document_version_field_value_id": 1,
		"master_file_id": 1,
		"previous_value": "null",
		"modified_value": "BN010011"
	},
	{
	"id": 2,
	"document_version_field_value_id": 2,
	"master_file_id": 2,
	"previous_value": null,
	"modified_value": "THLCH"
},
{
	"id": 3,
	"document_version_field_value_id": 3,
	"master_file_id": 3,
	"previous_value": null,
	"modified_value": "SGSIN"
},
{
	"id": 4,
	"document_version_field_value_id": 4,
	"master_file_id": 4,
	"previous_value": null,
	"modified_value": "BN020022"
},
{
	"id": 5,
	"document_version_field_value_id": 5,
	"master_file_id": 5,
	"previous_value": null,
	"modified_value": "SGSIN"
},
{
	"id": 6,
	"document_version_field_value_id": 6,
	"master_file_id": 6,
	"previous_value": null,
	"modified_value": "THLCH"
}
	
]`)

func SeedLogDocumentVersionFieldValue(db *gorm.DB) error {
	var data []models.LogDocumentVersionFieldValue
	seed := SEED_LOG_DOCUMENT_VERSION_FIELD_VALUE
	if err := BaseSeeder(db, &data, seed); err != nil {
		fmt.Println("------")
		fmt.Println(err)
		fmt.Println("------")
		return err
	}
	return nil
}
