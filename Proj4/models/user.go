package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname string
	Surname   string
}
