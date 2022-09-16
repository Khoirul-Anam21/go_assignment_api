package repositories

import (
	// "log"
	// "go_assignment_api/models"
	// "github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ItemDBRepository struct {
	DB gorm.DB
}

func GenerateItemRepository(DB gorm.DB) *ItemDBRepository{
	return &ItemDBRepository{DB: DB};
}
