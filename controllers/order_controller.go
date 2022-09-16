package controllers

import (
	"go_assignment_api/services"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	serviceProvider services.ServiceProvider
}

func (oc OrderController) RespondCreateOrder(c *gin.Context){
	response := oc.serviceProvider.AddOrder(c);
	c.JSON(response["status"].(int), response);
}

func (oc OrderController) RespondGetAllOrders(c *gin.Context){
	response := oc.serviceProvider.GetOrders(c);
	c.JSON(response["status"].(int), response);
}

func (oc OrderController) RespondUpdateOrder(c *gin.Context){
	response := oc.serviceProvider.UpdateOrder(c);
	c.JSON(response["status"].(int), response);
}

func (oc OrderController) RespondDeleteOrder(c *gin.Context){
	response := oc.serviceProvider.DeleteOrder(c);
	c.JSON(response["status"].(int), response);
}

func ProvideController(orderService services.ServiceProvider) *OrderController {
	return &OrderController{serviceProvider: orderService};
}

