package models

import "gorm.io/gorm"

// Section struct
type Section struct {
	gorm.Model
	Title           string
	Description     string
	SectionSchemaID int
	SectionSchema   SectionSchema
	PostID          int
	Post            Post
	Data            string
}
