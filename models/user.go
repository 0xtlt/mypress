package models

import "gorm.io/gorm"

// User struct
type User struct {
	gorm.Model
	Username    string
	Email       string
	Password    string
	Token       string
	Description string
	IsAdmin     bool
	AvatarID    int
	Avatar      Media
}
