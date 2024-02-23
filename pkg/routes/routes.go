package routes

import (
	"github.com/Ashutosh1921/BookStore/pkg/controller"
	"github.com/gorilla/mux"
)

var RegisterbookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controller.GetBook).Methods("GET")
	router.HandleFunc("/book/", controller.CreateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controller.GetBookid).Methods("GET")
	router.HandleFunc("/book/{bookId}", controller.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{bookId}", controller.UpdateBook).Methods("PUT")

}
