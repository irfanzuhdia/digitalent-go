package domain

import (
	"time"
)

type Photo struct {
	ID        uint   `gorm:"primaryKey" json:"id" form:"id"`
	Title     string `gorm:"not null" json:"title" form:"title"`
	Caption   string `json:"caption" form:"caption"`
	PhotoUrl  string `gorm:"not null" json:"photo_url" form:"photo_url"`
	UserID    uint   `gorm:"not null" json:"user_id" form:"user_id"`
	User      User
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}
