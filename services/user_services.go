package services

import (
	// "go_assignment_api/helpers"
	"fmt"
	"go_assignment_api/helpers"
	"go_assignment_api/models"
	"go_assignment_api/repositories"
	"log"
	"net/http"
)

type ServiceUserProvider struct {
	dbProvider repositories.UserDBRepository
}

func (sp ServiceUserProvider) ServiceAddUser(user models.User) (models.User, uint){
	// contentType := helpers.GetContentType()  taro di controller
	// err := json.NewDecoder(c.Request.Body).Decode(&user) taro di controller

	user, err := sp.dbProvider.AddUser(user)
	if err != nil {
		log.Println(err)
		return models.User{}, http.StatusBadRequest
	}

	return user, http.StatusCreated;

}

func (sp ServiceUserProvider) ServiceLogin(user models.User) (string, uint){
	passwordReq := user.Password;
	user, err := sp.dbProvider.GetUser(user);
	if err != nil {
		log.Println(err)
		return "", http.StatusUnauthorized
	}
	comparePass := helpers.CompareHash([]byte(user.Password), []byte(passwordReq));

	if !comparePass {
		log.Println(err)
		return "", http.StatusUnauthorized
	}

	fmt.Println(user)

	token := helpers.GenerateToken(user.ID, user.Email);
	return token, http.StatusOK

}


func ProvideUserService(userRepo repositories.UserDBRepository) *ServiceUserProvider {
	return &ServiceUserProvider{dbProvider: userRepo}
}
