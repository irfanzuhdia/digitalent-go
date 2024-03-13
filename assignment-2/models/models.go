package models

import (
	"time"
)

type Order struct {
	ID           uint      `gorm:"primaryKey" json:"lineItemId"`
	CustomerName string    `gorm:"not null; type:varchar(50)" json:"customerName"`
	OrderedAt    time.Time `gorm:"not null;" json:"orderedAt"`
	Items        []Item    `gorm:"foreignKey:OrderId" json:"items"`
}

type Item struct {
	ID          uint   `gorm:"primaryKey" json:"lineItemId"`
	Code        string `gorm:"not null; type:varchar(10)" json:"customerName"`
	Description string `gorm:"not null; type:varchar(50)" json:"description"`
	Quantity    uint   `gorm:"not null" json:"quantity"`
	OrderID     uint
}
