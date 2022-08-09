package config

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	

	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error", err)
		return nil
	}

	return db
}
