package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	models "github.com/Ashutosh1921/BookStore/pkg/Models"
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
