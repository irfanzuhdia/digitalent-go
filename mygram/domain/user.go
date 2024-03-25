package domain

import (
	"time"
)

type User struct {
	ID              uint      `gorm:"primaryKey" json:"id" form:"id"`
	Email           string    `gorm:"unique;not null" json:"email" form:"email"`
	Username        string    `gorm:"unique;not null" json:"username" form:"username"`
	Password        string    `gorm:"not null" json:"password" form:"password"`
	Age             int       `gorm:"not null" json:"age" form:"age"`
	ProfileImageUrl string    `json:"profile_image_url" form:"profile_image_url"`
	CreatedAt       time.Time `json:"created_at" form:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" form:"updated_at"`
}
