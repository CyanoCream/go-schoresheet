package models

import (
	"gorm.io/gorm"
	"time"
)

type Tbl_match struct {
	gorm.Model
	Club_id1       int       `json:"club_id1"`
	Club_id2       int       `json:"club_id2"`
	Date_match     time.Time `json:"date_match"`
	Level          int       `gorm:"not null" json:"level" valid:"required~Level is required"`
	Turnament_code int       `gorm:"not null" json:"turnament_code" valid:"required~Turnamen_code is required"`
}
