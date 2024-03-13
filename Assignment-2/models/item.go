package models

type Item struct {
	Id          uint   `gorm:"primaryKey" json:"id"`
	ItemCode    string `gorm:"not null; type:varchar(10)" json:"code"`
	Description string `gorm:"not null; type:varchar(50)" json:"description"`
	Quantity    uint   `gorm:"not null" json:"quantity"`
	OrderId     uint   `json:"order_id"`
}
