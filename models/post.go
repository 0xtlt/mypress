package models

import (
	"time"

	"gorm.io/gorm"
)

// Post struct
type Post struct {
	gorm.Model
	Title       string
	Description string
	IsPage      bool
	Handle      string
	TemplateID  int
	Template    Template
	PublishedAt time.Time
}
