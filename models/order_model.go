package models

import (
	"time"
)

type Order struct {
	OrderId int `gorm:"primaryKey;auto_increment" json:"orderId"`
	CustomerName string `json:"customerName"`
	OrderedAt time.Time `json:"orderedAt"`
	Items []Item `gorm:"foreignKey:OrderId" json:"items"`
}