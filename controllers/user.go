package controllers

import (
	"mypress/database"
	"mypress/models"

	"golang.org/x/crypto/bcrypt"
)

// CreateUser func
func CreateUser(user models.User) uint {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user.Password = string(hashedPassword)
	database.Get().Create(&user)
	return user.ID
}
