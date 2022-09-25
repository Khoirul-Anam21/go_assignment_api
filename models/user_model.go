package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string `json:"full_name" form:"full_name" valid:"required~Fullname required"`
	Email string `json:"email" form:"email" valid:"required~Email required"`
	Password string `json:"password" form:"password" valid:"required~Password required"`
	Products []Product `json:products`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error){
	_, errCreate := govalidator.ValidateStruct(u);
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}