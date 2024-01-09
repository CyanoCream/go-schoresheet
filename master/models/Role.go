package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Code  string `gorm:"not null;uniqueIndex" json:"code" valid:"required"`
	Name  string `gorm:"not null;uniqueIndex" json:"name" valid:"required"`
	Guard string `json:"guard" valid:"required"`
	Tag   string `json:"tag"`
}
