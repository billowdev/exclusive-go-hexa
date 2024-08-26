package seeders

import (
	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	"gorm.io/gorm"
)

var SEED_ORDER = []byte(`[
	{
		"order_type": "internal",
		"port_of_loading": "pol_1",
		"port_of_destination": "pod_1",
		"description_of_goods": "dog_1"
	},
		{
		"order_type": "internal",
		"port_of_loading": "pol_2",
		"port_of_destination": "pod_2",
		"description_of_goods": "dog_2"
	},
		{
		"order_type": "external",
		"port_of_loading": "pol_3",
		"port_of_destination": "pod_3",
		"description_of_goods": "dog_3"
	}
]`)

func SeedOrder(db *gorm.DB) error {
	var data []models.Order
	seed := SEED_ORDER
	if err := BaseSeeder(db, &data, seed); err != nil {
		return err
	}
	return nil
}
