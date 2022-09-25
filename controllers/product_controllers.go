package controllers

import (
	"go_assignment_api/helpers"
	"go_assignment_api/models"
	"go_assignment_api/services"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	serviceProductProvider services.ServiceProductProvider
}

func (pc ProductController) CreateProduct(c *gin.Context){
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Product := models.Product{}
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userId
	response, statusCode := pc.serviceProductProvider.ServiceAddProduct(Product)

	if statusCode == http.StatusBadRequest {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
			"message": "Invalid request bro",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"added product": response,
	})
	

}

func (pc ProductController) RespondUpdateProduct(c *gin.Context){
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	product := models.Product{}
	productId, _ := strconv.Atoi(c.Param("productId"));

	userId := uint(userData["id"].(float64))
	
	if contentType == appJSON {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}
	product.UserID = userId
	product.ID = uint(productId)

	response, statusCode := pc.serviceProductProvider.ServiceUpdateProduct(product, uint(productId))
	if statusCode == http.StatusInternalServerError {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "kesalahan sistem",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success update product",
		"data": response,
	})
}



func ProvideProductController(productService services.ServiceProductProvider) *ProductController {
	return &ProductController{serviceProductProvider: productService};
}