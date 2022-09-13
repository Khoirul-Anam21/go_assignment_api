package config

import (
	"fmt"
	"go_assignment_api/models"

	// "gorm.io/gorm"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/go_api_order?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=True&loc=local");
	if err != nil {
		fmt.Println(err)
		panic("failed to connect to database: ");
	}
	db.AutoMigrate(models.Order{});
	db.AutoMigrate(models.Item{});
	return db
}