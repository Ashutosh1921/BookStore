package routes

import (
	"github.com/Ashutosh1921/BookStore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterbookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/", controllers.CreateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.GetBookid).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")

}
