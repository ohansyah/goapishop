package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Token struct
type Token struct {
	gorm.Model
	DeviceID     string    `gorm:"type:varchar(50); NOT NULL" json:"device_id"`
	DeviceType   string    `gorm:"type:varchar(25);" json:"device_type"`
	TokenCode    string    `gorm:"type:varchar(256); NOT NULL index:token_code" json:"token_code"`
	RefreshToken string    `gorm:"type:varchar(256);" json:"refresh_token"`
	ExpiredDate  time.Time `gorm:"type:datetime; index:expired_date" json:"expired_date"`
}

// Tokens Arr
type Tokens []Token
