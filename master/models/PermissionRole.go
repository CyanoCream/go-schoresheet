package models

import "gorm.io/gorm"

type PermissionRole struct {
	gorm.Model
	RoleCode       string `gorm:"not null" json:"role_code" valid:"required"`
	PermissionCode string `gorm:"not null" json:"permission_code" valid:"required"`
}
