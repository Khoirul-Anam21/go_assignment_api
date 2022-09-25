package repositories

import (

	"go_assignment_api/models"
	// "io/ioutil"
	"log"

	// "github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"go_assignment_api/helpers"
)

type UserDBRepository struct {
	DB gorm.DB
}

func (dp UserDBRepository) AddUser(user models.User) (models.User, error) {
	// var (
	// 	user models.User
	// )
	user.Password = helpers.HashPass(user.Password);
	err := dp.DB.Create(&user).Error
	if err != nil {
		log.Println("ERROR -> Invalid SQL Syntax")
		return models.User{}, err;
	}
	return user, err
}

func (dp UserDBRepository) GetUser(user models.User) (models.User, error) {
	err := dp.DB.Debug().Where("email= ?", user.Email).Take(&user).Error
	if err != nil {
		log.Println("ERROR -> Invalid SQL Syntax")
		return models.User{}, err;
	}
	return user, err
}

func GenerateUserRepository(DB gorm.DB) *UserDBRepository {
	return &UserDBRepository{DB: DB}
}