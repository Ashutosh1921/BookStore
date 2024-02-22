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

func GetAllBook() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByid(id int64) (*Book, *gorm.DB) {
	var getbook Book
	db := db.Where("ID=?", id).Find(&getbook)
	return &getbook, db
}

func DeleteBook(id int64) Book {
	// var book Book
	// db.Find("ID=?", id).Delete(book)
	// return book
	// this is gpt code more safe
	var book Book
	db.Where("ID=?", id).First(&book)
	db.Delete(&book)
	return book

}
