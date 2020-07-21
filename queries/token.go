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

// CreateTokenLog data
func CreateTokenLog(data models.TokenLog) {
	db := database.ConnectToDB()
	db.Create(&data)
}

// GetTokenData get token by token string
func GetTokenData(tokenCode string) models.Token {
	var token models.Token
	db := database.ConnectToDB()
	db.Where("token_code = ?", tokenCode).Find(&token)
	return token
}

// GetTokenProfile get token profiles by token id
func GetTokenProfile(ID uint) models.TokenProfile {
	var tokenProfile models.TokenProfile
	db := database.ConnectToDB()
	db.Where("token_id = ?", ID).Find(&tokenProfile)
	return tokenProfile
}

// UpdateTokenProfile data
func UpdateTokenProfile(ID uint, tokenID uint, userID uint) *gorm.DB {
	db := database.ConnectToDB()
	var tokenProfile models.TokenProfile
	result := db.Model(tokenProfile).Where("id = ?", ID).Updates(&models.TokenProfile{TokenID: tokenID, UserID: userID, LastActivity: time.Now()})
	return result
}

// CreateTokenProfile data
func CreateTokenProfile(data models.TokenProfile) {
	db := database.ConnectToDB()
	db.Create(&data)
}
