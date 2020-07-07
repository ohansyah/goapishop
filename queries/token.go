package queries

import (
	"api_olshop/database"
	"api_olshop/models"
	"time"

	"github.com/jinzhu/gorm"
)

// GetTokenByDevID get token by device id
func GetTokenByDevID(deviceID string) models.Token {
	var token models.Token
	db := database.ConnectToDB()
	db.Where("device_id = ?", deviceID).Find(&token)
	return token
}

// CreateToken data
func CreateToken(deviceID string, deviceType string, tokenCode string, refreshToken string, expiredDate time.Time) models.Token {
	db := database.ConnectToDB()
	var result = &models.Token{DeviceID: deviceID, DeviceType: deviceType, TokenCode: tokenCode, RefreshToken: refreshToken, ExpiredDate: expiredDate}
	db.Create(&result)
	return *result
}

// UpdateToken data
func UpdateToken(ID uint, tokenCode string, refreshToken string, expiredDate time.Time) *gorm.DB {
	db := database.ConnectToDB()
	var token models.Token
	result := db.Model(token).Where("id = ?", ID).Updates(&models.Token{TokenCode: tokenCode, RefreshToken: refreshToken, ExpiredDate: expiredDate})
	return result
}
