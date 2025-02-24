package models

import (
	"github.com/asaskevich/govalidator"
	"go-blognews/helpers"
	"gorm.io/gorm"
)

type Users struct {
	GormModel
	Username string `gorm:"unique;not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Email    string `gorm:"unique;not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Your password must be at least 6 characters"`
	RoleId   uint   `gorm:"not null" json:"role_id" form:"role_id" valid:"required~Your role is required"`
	Roles    *Roles `gorm:"foreignKey:RoleId;references:ID" json:"roles"`
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
