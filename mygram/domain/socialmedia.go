package domain

import (
	"time"
)

type SocialMedia struct {
	ID             uint   `gorm:"primaryKey" json:"id" form:"id"`
	Name           string `gorm:"not null" json:"name" form:"name"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url"`
	UserID         uint   `gorm:"not null" json:"user_id" form:"user_id"`
	User           User
	CreatedAt      time.Time `json:"created_at" form:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" form:"updated_at"`
}
