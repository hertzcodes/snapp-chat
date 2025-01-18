package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Groups []Room `gorm:"many2many:user_groups"`
}


