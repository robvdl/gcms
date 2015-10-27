package db

import (
	"github.com/robvdl/gcms/models"
)

// GetUserByUsername does a query that returns the User matching that username.
func GetUserByUsername(username string) *models.User {
	var user models.User
	DB.Where("username = ?", username).First(&user)
	return &user
}
