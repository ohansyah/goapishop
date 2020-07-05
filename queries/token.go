package queries

import (
	"api_olshop/database"
	"api_olshop/models"
)

// ValidateTokenApp Get Single token apps
func ValidateTokenApp(Name string, SecretKey string) bool {
	db := database.ConnectToDB()
	var app models.App
	if result := db.Where("name = ? and secret_key = ?", Name, SecretKey).Find(&app); result.Error != nil {
		return false
	}
	return true
}
