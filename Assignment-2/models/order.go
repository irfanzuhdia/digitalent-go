package models

import (
	"time"
)

type Order struct {
	Id           uint      `gorm:"primaryKey" json:"orderId"`
	OrderedAt    time.Time `gorm:"not null;" json:"ordered_at"`
	CustomerName string    `gorm:"not null; type:varchar(50)" json:"customer_name"`
	Items        []Item    `gorm:"foreignKey:order_id" json:"items"`
}
