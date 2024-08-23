package models

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID        string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password"`
	FirstName string    `json:"first_name" gorm:"size:255"`
	LastName  string    `json:"last_name" gorm:"size:255"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
