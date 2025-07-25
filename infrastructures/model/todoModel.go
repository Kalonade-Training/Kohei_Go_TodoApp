package model

import (
	"time"
)
//gormでテーブル作成に使うmodel
type Todo struct {
	ID string `gorm:"type:char(36);primaryKey"`
	UserID string `gorm:"type:char(36);not null"`
	Title string`gorm:"type:char(50);not null"`
	Description *string `gorm:"type:text"`
	DueDate *time.Time `gorm:"type:date"`
	CompletedAt *time.Time `gorm:"type:date"`
	Status string `gorm:"type:char(15);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}