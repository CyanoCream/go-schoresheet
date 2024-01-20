package models

import "gorm.io/gorm"

type Tbl_player struct {
	gorm.Model
	Name      string `gorm:"not null" json:"name" valid:"required-Name is required"`
	Club_id   int    `json:"club_id"`
	Position  string `json:"position"`
	Is_active bool   `gorm:"not null;default:true" json:"is_active"`
}
