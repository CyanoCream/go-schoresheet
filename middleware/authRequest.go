package middleware

import (
	"gorm.io/gorm"
	"time"
)

type LoginField struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type JWT struct {
	Token string `json:"token"`
}

type Session struct {
	gorm.Model
	UserID     int       `json:"user_id"`
	Token      string    `json:"token"`
	LoginTime  time.Time `json:"login_time"`
	LogoutTime time.Time `json:"logout_time"`
}
