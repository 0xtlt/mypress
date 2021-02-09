package models

import (
	"gorm.io/gorm"
)

// Theme struct
type Theme struct {
	gorm.Model
	Name string
}
