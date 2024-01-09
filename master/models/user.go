package models

import (
	"github.com/asaskevich/govalidator"
	"go-scoresheet/master/helpers"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname string `gorm:"not null" json:"fullname" valid:"required~Fullname is required"`
	Username string `gorm:"not null;uniqueIndex" json:"username" valid:"required~Username is required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" valid:"required~Email is required,email~Invalid email format"`
	Password string `gorm:"not null" json:"password" valid:"required~Password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return
	}

	hashedPass, err := helpers.HashPassword(u.Password)
	if err != nil {
		return
	}
	u.Password = hashedPass

	return
}
