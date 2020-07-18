package models

import (
	"github.com/jinzhu/gorm"
)

// Role struct
type Role struct {
	gorm.Model
	Name   string `gorm:"type:varchar(50);" json:"name"`
	Desc   string `gorm:"type:text;" json:"desc"`
	Status string `gorm:"type:tinyint; index:status" json:"status"`
}

// Roles Arr
type Roles []Role
