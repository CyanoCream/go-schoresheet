package models

import "gorm.io/gorm"

type Tbl_score struct {
	gorm.Model
	Match_id int    `gorm:"not null" json:"match_id" valid:"required~Match_id is required"`
	Club_id  int    `gorm:"not null" json:"club_id" valid:"required~Club_id is required"`
	Score    int    `json:"score"`
	Status   string `json:"status"`
}
