package models

import (
	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model
	UserId    uint
	ProductId uint
}
