package queries

import (
	"api_olshop/database"
	"api_olshop/models"
)

// GetTokenApp Get Single token apps
func GetTokenApp(Name string, SecretKey string) models.App {
	db := database.ConnectToDB()
	var app models.App
	db.Where("name = ? and secret_key = ?", Name, SecretKey).Find(&app)
	return app
}
