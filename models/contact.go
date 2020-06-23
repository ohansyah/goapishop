package models

import "github.com/jinzhu/gorm"

// Contact struct
type Contact struct {
	gorm.Model
	Name    string `gorm:"type:varchar(255); NOT NULL" json:"name" binding:"required"`
	Email   string `gorm:"type:varchar(255);" json:"email"`
	Phone   string `gorm:"type:varchar(100); NOT NULL; UNIQUE; UNIQUE_INDEX" json:"phone" binding:"required"`
	Address string `gorm:"type:text" json:"address"`
}

// Contacts Arr
type Contacts []Contact
