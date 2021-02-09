package models

import "gorm.io/gorm"

// Redirection struct
type Redirection struct {
	gorm.Model
	Handle string
	To     string
}
