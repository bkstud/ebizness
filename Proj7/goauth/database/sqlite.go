package database

import (
	"goauth/structs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB = nil

func init() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&structs.Session{})
	database = db

}

func GetDatabase() *gorm.DB {
	return database
}
