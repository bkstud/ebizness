package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Firstname string
	Surname   string
}
