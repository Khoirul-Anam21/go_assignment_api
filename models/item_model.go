package models


type Item struct {
	ItemId int `gorm:"primaryKey;auto_increment" json:"itemId"`
	ItemCode string `json:"itemCode"`
	Description string `json:"description"`
	Quantity int `json:"quantity"`
	OrderId uint `json:"orderId"`
}