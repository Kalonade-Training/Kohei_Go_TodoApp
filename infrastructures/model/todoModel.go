package model

import (
	"time"
)
//gormでテーブル作成に使うmodel
type Todo struct {
	ID string `gorm:"type:varchar(36);primaryKey"`
	UserID string `gorm:"type:varchar(36);not null"`
	Title string`gorm:"type:varchar(50);not null"`
	Description *string `gorm:"type:text"`
	DueDate *time.Time `gorm:"type:date"`
	CompletedAt *time.Time `gorm:"type:date"`
	Status string `gorm:"type:varchar(15);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}