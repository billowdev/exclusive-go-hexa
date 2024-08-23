package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderModel struct {
	gorm.Model
	ID                 string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	PortOfLoading      string    `json:"port_of_loading"`
	PortOfDestination  string    `json:"port_of_destination"`
	DescriptionOfGoods string    `json:"description_of_goods"`
	CreatedAt          time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

var TNOrder = "orders"

func (st *OrderModel) TableName() string {
	return TNOrder
}
