package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Code        string `gorm:"type:varchar(255);not null" json:"code"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	GuardName   string `gorm:"type:varchar(255);not null" json:"guard_name"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Action      string `gorm:"type:char(1);not null" json:"action"`
	Module      string `gorm:"type:varchar(255);not null" json:"module"`
	Tag         string `gorm:"type:varchar(255);not null" json:"tag"`
	IsActive    bool   `gorm:"not null" json:"is_active"`
}
