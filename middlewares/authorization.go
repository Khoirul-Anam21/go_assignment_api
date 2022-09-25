package middlewares

import (
	"go_assignment_api/models"
	// "go_assignment_api/repositories"
	"net/http"
	"strconv"

	// "net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductAuthorization(tx *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		productId, err := strconv.Atoi(c.Param("productId"));
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
				"message": "Invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64));
		product := models.Product{}

		err = tx.Select("user_id").First(&product, uint(productId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "data not found",
				"message": "data doen't mencet bel",
			})
			return
		}

		if product.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
				"message": "ga boleh masuk bro",
			})
			return
		}
		

		// err = repositories.SelectProduct()

	}
}