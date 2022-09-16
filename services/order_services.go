package services

import (
	"encoding/json"
	"fmt"
	"go_assignment_api/models"
	"go_assignment_api/repositories"
	// "io/ioutil"

	// "log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ServiceProvider struct {
	dbProvider repositories.OrderDBRepository
}

func (sp ServiceProvider) AddOrder(c *gin.Context) gin.H {
	var (
		order models.Order
	)
	err := json.NewDecoder(c.Request.Body).Decode(&order)
	if err != nil {
		return gin.H{
			"status":  http.StatusBadRequest,
			"message": "invalid body request",
		}
	}
	order, err = sp.dbProvider.AddOrder(order)
	if err != nil {
		return gin.H{
			"status":  http.StatusBadRequest,
			"message": "Bad request",
		}
	}
	return gin.H{
		"status":  http.StatusCreated,
		"message": "Success add new order",
	}

}

func (sp ServiceProvider) GetOrders(c *gin.Context) gin.H {
	orders, err := sp.dbProvider.GetOrders()
	if err != nil {
		return gin.H{
			"status":  http.StatusNotFound,
			"message": "Bad request",
		}
	}
	return gin.H{
		"status": http.StatusOK,
		"orders": orders,
	}
}

func (sp ServiceProvider) UpdateOrder(c *gin.Context) gin.H {
	var (
		order models.Order
	)
	err := json.NewDecoder(c.Request.Body).Decode(&order)
	// fmt.Printf("%+v\n", order)
	// fmt.Println(ioutil.ReadAll(c.Request.Body))
	if err != nil {
		fmt.Println(err)
		return gin.H{
			"status":  http.StatusBadRequest,
			"message": "invalid body request",
		}
	}
	order, err = sp.dbProvider.UpdateOrders(order, strconv.Itoa(order.OrderId))
	if err != nil {
		return gin.H{
			"status":  http.StatusNotFound,
			"message": "data not found",
		}
	}
	return gin.H{
		"status":           http.StatusOK,
		"order updated to": order,
	}
}

func (sp ServiceProvider) DeleteOrder(c *gin.Context) gin.H{
	id := c.Param("id");
	msg, err := sp.dbProvider.DeleteOrders(id);
	if err != nil {
		return gin.H{
			"status": http.StatusNotFound,
			"message": "order not found",
		}
	}
	return gin.H{
		"status": http.StatusOK,
		"message": msg,
	}
}

func ProvideService(orderRepo repositories.OrderDBRepository) *ServiceProvider {
	return &ServiceProvider{dbProvider: orderRepo}
}
