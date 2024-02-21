package main

import (
	"log"
	"net/http"

	routes "github.com/Ashutosh1921/BookStore/pkg/Routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterbookstoreRoutes(r)
	// Now creating a handle APIs
	http.Handle("/", r)
	// now initializing server here
	log.Fatal(http.ListenAndServe(":40000", r))

}
