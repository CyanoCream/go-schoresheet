package models

import "gorm.io/gorm"

type Turnament struct {
	gorm.Model
	Name string `gorm:"not null" json:"name" valid:"required~Turnament is required"`
}
