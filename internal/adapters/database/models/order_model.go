package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID                 uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderType          string         `gorm:"orderType" json:"orderType"`
	CreatedAt          time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	PortOfLoading      string         `json:"port_of_loading"`
	PortOfDestination  string         `json:"port_of_destination"`
	DescriptionOfGoods string         `json:"description_of_goods"`
}

var TNOrder = "orders"

func (st *Order) TableName() string {
	return TNOrder
}
