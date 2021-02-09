package controllers

import (
	"log"
	"mypress/database"
	"mypress/models"
)

// CheckWebsite func
func CheckWebsite() bool {
	var website models.Website
	result := database.Get().Limit(1).Find(&website)

	return result.RowsAffected == 1
}

// GetWebsite func
func GetWebsite() models.Website {
	var website models.Website
	result := database.Get().Limit(1).Find(&website)

	if result.RowsAffected == 0 {
		log.Fatal("Get Website returns 0 row")
	}

	return website
}
