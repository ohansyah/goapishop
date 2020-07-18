package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// TokenProfile struct
type TokenProfile struct {
	gorm.Model
	TokenID      int        `gorm:"type:int; NOT NULL index:token_id" json:"token_id"`
	UserID       int        `gorm:"type:int; NOT NULL index:user_id" json:"user_id"`
	LastActivity *time.Time `gorm:"type:datetime; json:"last_activity"`
}

// TokenProfiles Arr
type TokenProfiles []TokenProfile
