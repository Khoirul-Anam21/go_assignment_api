package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	item_code string
	description string
	quantity int
	order_id uint
}