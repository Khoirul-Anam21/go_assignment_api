package router

import (
	"go_assignment_api/controllers"
	"go_assignment_api/middlewares"

	"github.com/gin-gonic/gin"
)

func RunOrderRouter(router *gin.Engine, controllers *controllers.OrderController) {
	router.POST("/orders", controllers.RespondCreateOrder)
	router.GET("/orders", controllers.RespondGetAllOrders)
	router.PUT("/orders", controllers.RespondUpdateOrder)
	router.DELETE("/orders/:id", controllers.RespondDeleteOrder)
}

func RunUserRouter(router *gin.Engine, controllers *controllers.UserController) {

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllers.RespondCreateUser);
		userRouter.POST("/login", controllers.RespondUserLogin);
	}
}

func RunProductRouter(router *gin.Engine, controllers *controllers.ProductController){
	productRouter := router.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:productId", controllers.RespondUpdateProduct)
	}
}
