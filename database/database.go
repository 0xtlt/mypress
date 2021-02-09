package database

import (
	"fmt"
	"mypress/models"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Get func
func Get() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("mypress.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}

	db.AutoMigrate(
		&models.Theme{},
		&models.Template{},
		&models.Website{},
		&models.User{},
		&models.Post{},
		&models.Media{},
		&models.Redirection{},
		&models.SectionSchema{},
		&models.Section{},
	)
	return db
}
