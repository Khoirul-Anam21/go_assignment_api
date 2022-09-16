package router

import (
	"go_assignment_api/controllers"

	"github.com/gin-gonic/gin"
)

func RunRouter(router *gin.Engine, controllers *controllers.OrderController) {
	router.POST("/orders", controllers.RespondCreateOrder);
	router.GET("/orders", controllers.RespondGetAllOrders);
	router.PUT("/orders", controllers.RespondUpdateOrder);
	router.DELETE("/orders/:id", controllers.RespondDeleteOrder)
}