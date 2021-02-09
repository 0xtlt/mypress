package models

import "gorm.io/gorm"

// Media struct
type Media struct {
	gorm.Model
	Format string
	Name   string
	Size   int64
}
