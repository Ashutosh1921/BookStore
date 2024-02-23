package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	models "github.com/Ashutosh1921/BookStore/pkg/Models"
	utils "github.com/Ashutosh1921/BookStore/pkg/Utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBook()
	res, err := json.Marshal(newBooks)
	if err != nil {
		log.Fatal(err)
	}
	// Now setting the header type
	w.Header().Set("Content-type", "application/json")
	// here setting what type of value does the response contains
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	// this will send res to frontend
}

func GetBookid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	// finding a bookid variable and assign it to the bookId
	ID, err := strconv.ParseInt(bookId, 0, 0)
	// here changing string to int with automatic base and automatic bit size by filling 0,0
	if err != nil {
		fmt.Println("error while parsing")
	}
	// Now we will fetch the book by giving id to db
	bookDetails, _ := models.GetBookByid(ID)
	// how function is models package will fetch the details from database
	res, _ := json.Marshal(bookDetails)
	// now will send the fetched data to frontend by converting it json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	// res will get converted into http response
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := models.Book{}
	utils.ParseBody(r, createBook)
	// Now parsing the request body into empty book struct
	b := createBook.CreateBook()
	// here we are calling CreateBook a receiver function of type Book
	// which create a this new book inside the database.
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	// finding a bookid variable and assign it to the bookId
	ID, err := strconv.ParseInt(bookId, 0, 0)
	// here changing string to int with automatic base and automatic bit size by filling 0,0
	if err != nil {
		fmt.Println("error while parsing")
	}
	// Now that we got the id of the book we delete that from the database
	book := models.DeleteBook(ID)
	// Now that book is deleted and also got details of book
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	// gettting address of empty book
	utils.ParseBody(r, updateBook)
	// parsing the json request into the empty book variable
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	// here changing string to int with automatic base and automatic bit size by filling 0,0
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookdetails, db := models.GetBookByid(ID)
	if updateBook.Name != "" {
		bookdetails.Name = updateBook.Name
		// Now we are filling new data we got from the updatbook
	}
	if updateBook.Author != "" {
		bookdetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookdetails.Publication = updateBook.Publication
	}
	// Now we have updated the bookdetail we are saving it to database
	db.Save(&bookdetails)
	// db.Save is used to update the database if anything is updated.
	res, _ := json.Marshal(bookdetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
