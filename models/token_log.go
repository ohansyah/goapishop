package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// TokenLog struct
type TokenLog struct {
	gorm.Model
	TokenID    int        `gorm:"type:int; NOT NULL index:token_id" json:"token_id"`
	UserAgent  string     `gorm:"type:varchar(50);" json:"user_agent"`
	Path       string     `gorm:"type:varchar(256);" json:"path"`
	Method     string     `gorm:"type:varchar(10); NOT NULL json:"method"`
	Request    string     `gorm:"type:text;" json:"request"`
	Response   string     `gorm:"type:text;" json:"response"`
	Status     string     `gorm:"type:varchar(10);" json:"status"`
	APIVersion string     `gorm:"type:varchar(15); NOT NULL index:api_version"" json:"api_version"`
	StartDate  *time.Time `gorm:"type:datetime; json:"start_date"`
	EndDate    *time.Time `gorm:"type:datetime; json:"end_date"`
}

// TokenLogs Arr
type TokenLogs []TokenLog
