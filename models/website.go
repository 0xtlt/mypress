package models

import (
	"gorm.io/gorm"
)

// Website struct
type Website struct {
	gorm.Model
	Title       string
	Description string
	URL         string
	Font        string
	LogoID      int
	Logo        Media
	BannerID    int
	Banner      Media
	ThemeID     int
	Theme       Theme
}
