package services

import (
	// "go_assignment_api/helpers"
	// "fmt"
	// "go_assignment_api/helpers"
	"go_assignment_api/models"
	"go_assignment_api/repositories"
	"log"
	"net/http"
	// "strconv"
)

type ServiceProductProvider struct {
	dbProvider repositories.ProductDBRepository
}

func (sp ServiceProductProvider) ServiceAddProduct(product models.Product) (models.Product, uint){
	// contentType := helpers.GetContentType()  taro di controller
	// err := json.NewDecoder(c.Request.Body).Decode(&product) taro di controller

	product, err := sp.dbProvider.AddProduct(product)
	if err != nil {
		log.Println(err)
		return models.Product{}, http.StatusBadRequest
	}

	return product, http.StatusCreated;

}

func (sp ServiceProductProvider) ServiceUpdateProduct(product models.Product, productId uint) (models.Product, uint){
	// contentType := helpers.GetContentType()  taro di controller
	// err := json.NewDecoder(c.Request.Body).Decode(&product) taro di controller
	product, err := sp.dbProvider.UpdateProduct(product, productId);
	if err != nil {
		log.Println(err)
		return models.Product{}, http.StatusInternalServerError
	}

	return product, http.StatusOK;

}

func ProvideProductService(productRepo repositories.ProductDBRepository) *ServiceProductProvider {
	return &ServiceProductProvider{dbProvider: productRepo}
}
