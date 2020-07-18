package models

import (
	"github.com/jinzhu/gorm"
)

// User struct
type User struct {
	gorm.Model
	Name    string `gorm:"type:varchar(50);" json:"name"`
	Address string `gorm:"type:text;" json:"address"`
	RoleID  string `gorm:"type:int; NOT NULL" json:"role_id"`
	Phone   string `gorm:"type:varchar(15); index:phone; unique; NOT NULL" json:"phone"`
	Email   string `gorm:"type:varchar(50); index:email; unique; NOT NULL" json:"email"`
	Status  string `gorm:"type:tinyint; index:status" json:"status"`
}

// Users Arr
type Users []User
