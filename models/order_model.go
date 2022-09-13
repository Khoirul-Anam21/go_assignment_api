package models

import (
	"time"
)

type Order struct {
	order_id uint
	customer_name string
	orderedAt time.Time
	items []Item `gorm:"foreignKey:UserRefer"`
}