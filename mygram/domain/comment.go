package domain

import (
	"time"
)

type Comment struct {
	ID        uint   `gorm:"primaryKey" json:"id" form:"id"`
	Message   string `gorm:"not null" json:"message" form:"message"`
	PhotoID   uint   `gorm:"not null" json:"photo_id" form:"photo_id"`
	Photo     Photo  `gorm:"foreignKey:PhotoID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	UserID    uint   `gorm:"not null" json:"user_id" form:"user_id"`
	User      User
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}
