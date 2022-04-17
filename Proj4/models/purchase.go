package models

import (
	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model
	ProductID  uint
	Product    Product
	CustomerID uint
	Customer   Customer
}
