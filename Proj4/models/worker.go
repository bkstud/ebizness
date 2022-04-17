package models

import "gorm.io/gorm"

type Worker struct {
	gorm.Model
	Firstname string
	Surname   string
	Position  string
}
