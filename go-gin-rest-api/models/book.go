package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database!")
	}

	database.AutoMigrate(&Book{})

	DB = database
}
