package models

import "gorm.io/gorm"

type Tbl_club struct {
	gorm.Model
	Name      string `gorm:"not null" json:"name" valid:"required~Name is required"`
	Hometown  string `gorm:"not null" json:"hometown" valid:"required-Hometown is required"`
	Is_active bool   `gorm:"not null;default:true" json:"is_active"`
}
