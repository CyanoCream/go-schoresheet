package models

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	RoleCode string `gorm:"size:255;unique;not null" json:"role_code" binding:"required"`
	UserID   int    `gorm:"not null" json:"user_id" binding:"required"`
}
