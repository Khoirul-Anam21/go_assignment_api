package repositories

import (

	"go_assignment_api/models"
	// "io/ioutil"
	"log"

	// "github.com/gin-gonic/gin"
	"gorm.io/gorm"
	// "go_assignment_api/helpers"
)

type ProductDBRepository struct {
	DB gorm.DB
}

func (pr ProductDBRepository) AddProduct(product models.Product) (models.Product, error){
	err := pr.DB.Debug().Create(&product).Error
	if err != nil {
		log.Println(err)
		return models.Product{}, err;
	}
	return product, err
}

func (pr ProductDBRepository) SelectProduct(product models.Product, productId float64) (models.Product, error){
	err := pr.DB.Select("user_id").First(&product, uint(productId)).Error

	if err != nil {
		log.Println(err)
		return models.Product{}, err;
	}
	return product, err
}

func (pr ProductDBRepository) UpdateProduct(product models.Product, productId uint) (models.Product, error){
	err := pr.DB.Model(&product).Where("id = ?", productId).Updates(models.Product{Title: product.Title, Description: product.Description}).Error

	if err != nil {
		log.Println(err)
		return models.Product{}, err;
	}
	return product, err
}

func GenerateProductRepository(DB gorm.DB) *ProductDBRepository {
	return &ProductDBRepository{DB: DB}
}