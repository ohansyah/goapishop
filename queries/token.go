package queries

import (
	"api_olshop/database"
	"api_olshop/models"

	"github.com/jinzhu/gorm"
)

// GetTokenApp Get Single token apps
func GetTokenApp(Name string, SecretKey string) *gorm.DB {
	db := database.ConnectToDB()
	var app models.App
	return db.Where("name = ? and secret_key = ?", Name, SecretKey).Find(&app)
}
