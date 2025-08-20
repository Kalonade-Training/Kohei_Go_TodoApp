package model

import (
	"time"
)

type User struct {
	ID             string    `gorm:"type:char(36);primaryKey"`
	Username       string    `gorm:"type:varchar(50);unique;not null"`
	HashedPassword string    `gorm:"type:varchar(255);not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}