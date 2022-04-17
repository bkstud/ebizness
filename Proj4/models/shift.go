package models

import "gorm.io/gorm"

type Shift struct {
	gorm.Model
	WorkerID  uint
	Worker    Worker
	EnterTime string
	ExitTime  string
}
