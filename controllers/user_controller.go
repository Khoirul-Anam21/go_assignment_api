package controllers

import (
	"go_assignment_api/helpers"
	"go_assignment_api/models"
	"go_assignment_api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	serviceUserProvider services.ServiceUserProvider
}

var (
	appJSON = "application/json"
)

func (uc UserController) RespondCreateUser(c *gin.Context){

	contentType := helpers.GetContentType(c);
	user := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user);
	}
	response, statusCode := uc.serviceUserProvider.ServiceAddUser(user)

	if statusCode == http.StatusBadRequest {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request bro",
			"message": "Invalid request",
		})
		return;
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": response.ID,
		"user_name": response.FullName,
		"email": response.Email,
	})
}


func (uc UserController) RespondUserLogin(c *gin.Context) {
	contentType := helpers.GetContentType(c);
	user := models.User{}
	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user);
	}

	response, statusCode := uc.serviceUserProvider.ServiceLogin(user)

	if statusCode == http.StatusUnauthorized {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
			"message": "invalid email/password",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": response,
	})

}


func ProvideUserController(userService services.ServiceUserProvider) *UserController {
	return &UserController{serviceUserProvider: userService};
}