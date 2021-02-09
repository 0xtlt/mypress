package models

import "gorm.io/gorm"

// SectionSchema struct
type SectionSchema struct {
	gorm.Model
	Name        string
	Description string
	HTML        string
	CSS         string
	Javascript  string
	Schema      string // json format
}
