package repositories

import (
	// "fmt"
	"fmt"
	"go_assignment_api/models"
	// "io/ioutil"
	"log"

	// "github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderDBRepository struct {
	DB gorm.DB
}

func (dp OrderDBRepository) GetOrders() ([]models.Order, error) {
	var (
		orders []models.Order
	)
	err := dp.DB.Preload("Items").Find(&orders).Error
	if err != nil {
		log.Fatal("ERROR -> Invalid SQL Syntax")
	}
	return orders, err
}

func (dp OrderDBRepository) DeleteOrders(id string) (string, error) {
	var (
		order models.Order
		items []models.Item
	)
	// id := c.Param("id");
	err := dp.DB.First(&order, id).Error
	if err != nil {
		log.Println("Data not found")
	}
	err = dp.DB.Find(&items).Error
	fmt.Printf("%+v\n", order)
	if err != nil {
		log.Println("Data not found")
	}
	err = dp.DB.Delete(&items, "order_id LIKE ?", id).Error
	if err != nil {
		log.Println("failed to delete")
	}
	err = dp.DB.Delete(&order, order.OrderId).Error
	if err != nil {
		log.Println("failed to delete")
	}
	return "success delete data", err
}

func (dp OrderDBRepository) AddOrder(order models.Order) (models.Order, error) {
	// var (
	// 	order models.Order
	// )

	err := dp.DB.Create(&order).Error
	if err != nil {
		log.Println("ERROR -> Invalid SQL Syntax")

	}
	return order, err
}

func (dp OrderDBRepository) UpdateOrders(order models.Order, id string) (models.Order, error) {

	var (
		oldOrder models.Order
	)
	// tes, err := ioutil.ReadAll(c.Request.Body);
	err := dp.DB.Preload("Items").First(&oldOrder, id).Error
	if err != nil {
		log.Println("ERROR -> Invalid SQL Syntax")
	}

	err = dp.DB.Preload("Items").Model(&oldOrder).Updates(order).Error
	if err != nil {
		log.Println("ERROR -> Invalid SQL Syntax")
	}
	return order, err
}

func GenerateOrderRepository(DB gorm.DB) *OrderDBRepository {
	return &OrderDBRepository{DB: DB}
}
