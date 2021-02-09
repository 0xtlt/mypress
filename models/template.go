package models

import (
	"gorm.io/gorm"
)

// Template struct
type Template struct {
	gorm.Model
	Name    string
	Type    string
	ThemeID int
	Theme   Theme
	Code    string
}
