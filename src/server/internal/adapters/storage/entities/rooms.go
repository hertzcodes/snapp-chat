package entities

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Name  string
	Code  string
	Users []User `gorm:"many2many:user_groups"`
}
