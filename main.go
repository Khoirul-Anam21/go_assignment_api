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
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	routerEngine := gin.Default()
	db := conf.DBInit();
	repository := repositories.GenerateOrderRepository(*db);
	service := services.ProvideService(*repository);
	controllers := controllers.ProvideController(*service)

	router.RunRouter(routerEngine, controllers)
	routerEngine.Run(os.Getenv("PORT"))
	fmt.Println("Server berjalan pada http://localhost:3000")
}
