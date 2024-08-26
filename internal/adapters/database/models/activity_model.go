package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	// ID         string         `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	ID         uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	WorkflowID string         `gorm:"type:uuid;not null" json:"workflow_id"`
	Name       string         `gorm:"type:varchar(255);not null" json:"name"`
	Status     string         `gorm:"type:varchar(50);not null" json:"status"`
	Input      JSONB          `gorm:"type:jsonb" json:"input"`
	Output     JSONB          `gorm:"type:jsonb" json:"output"`
	Error      string         `gorm:"type:text" json:"error"`
	StartTime  time.Time      `json:"start_time"`
	EndTime    time.Time      `json:"end_time"`
}

var TNActivity = "activities"

func (st *Activity) TableName() string {
	return TNActivity
}