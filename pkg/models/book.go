package models

import (
	config "github.com/Ashutosh1921/BookStore/pkg/Config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm:"" json:"name"`
	// this is like key value pair here gorm represent the
	/*
		you can use struct tags to provide additional information
		about how a struct should be mapped to a database table or how
		it should be serialized to JSON.
	*/
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// Now we will start storing book data in database
func (b *Book) CreateBook() *Book {
	// creating new record
	db.NewRecord(b)
	db.Create(b)
	return b
}
