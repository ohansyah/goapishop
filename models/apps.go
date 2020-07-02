package models

import "github.com/jinzhu/gorm"

// App struct
type App struct {
	gorm.Model
	Name      string `gorm:"type:varchar(25); NOT NULL index:expired_date" json:"name"`
	SecretKey string `gorm:"type:varchar(50); NOT NULL" json:"secret_key"`
}

// Apps Arr
type Apps []App
