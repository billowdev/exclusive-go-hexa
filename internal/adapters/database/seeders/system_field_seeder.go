package seeders

import (
	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	"gorm.io/gorm"
)

var SEED_SYSTEM_FIELD = []byte(`[
	{
		"id": 1,
		"field_code": "FS01",
		"field_name": "Booking Number",
        "data_type": "text",
        "description": "Description for Field 1"
	},
		{
		"id": 2,
		"field_code": "FS02",
		"field_name": "Port Of Lading",
        "data_type": "text",
        "description": "Description for Field 2"
	},
	{
		"id": 3,
		"field_code": "FS03",
		"field_name": "Port Of Discharge",
        "data_type": "text",
        "description": "Description for Field 3"
	}
]`)

func SeedSystemField(db *gorm.DB) error {
	var data []models.SystemField
	seed := SEED_SYSTEM_FIELD
	if err := BaseSeeder(db, &data, seed); err != nil {
		return err
	}
	return nil
}

var SEED_GROUP_FIELD = []byte(`[
	{
	"id":1,
	"name": "shipment",
		"description": "Description for shipment"
	},
		{
	"id":2,
		"name": "document",
		"description": "Description for document"
	}
]`)

func SeedGroupField(db *gorm.DB) error {
	var data []models.SystemGroupField
	seed := SEED_GROUP_FIELD
	if err := BaseSeeder(db, &data, seed); err != nil {
		return err
	}
	return nil
}

var SEED_CONFIG_SYSTEM_MASTER_FILE_FIELD = []byte(`[
	{
		"system_field_id": 1,
		"system_group_field_id": 2
	},
	{
		"system_field_id": 2,
		"system_group_field_id": 1
	},
		{
		"system_field_id": 3,
		"system_group_field_id": 1
	}
]`)

func SeedConfigSystemMasterFileField(db *gorm.DB) error {
	var data []models.ConfigSystemMasterFileField
	seed := SEED_CONFIG_SYSTEM_MASTER_FILE_FIELD
	if err := BaseSeeder(db, &data, seed); err != nil {
		return err
	}
	return nil
}
