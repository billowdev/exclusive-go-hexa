package seeders

import (
	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	"gorm.io/gorm"
)

var SEED_MASTER_FILE = []byte(`[
	{
		"id": 1,
		"order_id": 1,
		"system_field_id": 1,
		"value": "BN010011"
	},
	{
		"id": 2,
		"order_id": 1,
		"system_field_id": 2,
		"value": "THLCH"
	},
	{
		"id": 3,
		"order_id": 1,
		"system_field_id": 3,
		"value": "SGSIN"
	},
	{
		"id": 4,
		"order_id": 2,
		"system_field_id": 1,
		"value": "BN020022"
	},
	{
		"id": 5,
		"order_id": 2,
		"system_field_id": 2,
		"value": "SGSIN"
	},
	{
		"id": 6,
		"order_id": 2,
		"system_field_id": 3,
		"value": "THLCH"
	}
]`)

func SeedMasterFile(db *gorm.DB) error {
	var data []models.MasterFile
	seed := SEED_MASTER_FILE
	if err := BaseSeeder(db, &data, seed); err != nil {
		return err
	}
	return nil
}

var SEED_LOG_MASTER_FILE = []byte(`[
	{
		"master_file_id": 1,
		"previous_value": "",
		"modified_value": "BN010011"
	},
	{
		"master_file_id": 2,
		"previous_value": "SGSIN",
		"modified_value": "THLCH"
	},
	{
		"master_file_id": 3,
		"previous_value": "THLCH",
		"modified_value": "SGSIN"
	},
		{
		"master_file_id": 4,
		"previous_value": "",
		"modified_value": "BN010011"
	},
	{
		"master_file_id": 5,
		"previous_value": "SGSIN",
		"modified_value": "THLCH"
	},
	{
		"master_file_id": 6,
		"previous_value": "THLCH",
		"modified_value": "SGSIN"
	}
	
]`)

func SeedLogMasterFile(db *gorm.DB) error {
	var data []models.LogMasterFile
	seed := SEED_LOG_MASTER_FILE
	if err := BaseSeeder(db, &data, seed); err != nil {
		return err
	}
	return nil
}
