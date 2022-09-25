package main

import (
	"fmt"
	conf "go_assignment_api/configs"
	"go_assignment_api/controllers"
	"go_assignment_api/repositories"
	"go_assignment_api/router"
	"go_assignment_api/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"
	// "gorm.io/gorm"
	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	routerEngine := gin.Default()
	db := conf.DBInit();
	// mainDB = db
	orderRepository := repositories.GenerateOrderRepository(*db);
	userRepository := repositories.GenerateUserRepository(*db);
	productRepository := repositories.GenerateProductRepository(*db);

	orderService := services.ProvideService(*orderRepository);
	userService := services.ProvideUserService(*userRepository);
	productService := services.ProvideProductService(*productRepository);

	orderControllers := controllers.ProvideController(*orderService);
	userControllers := controllers.ProvideUserController(*userService);
	productControllers := controllers.ProvideProductController(*productService);

	router.RunOrderRouter(routerEngine, orderControllers)
	router.RunUserRouter(routerEngine, userControllers)
	router.RunProductRouter(routerEngine, productControllers)

	routerEngine.Run(os.Getenv("PORT"))
	fmt.Println("Server berjalan pada http://localhost:3000")
}
