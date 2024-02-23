package main

import (
	"log"
	"net/http"

	routes "github.com/Ashutosh1921/BookStore/pkg/Routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterbookstoreRoutes(r)
	// Now creating a handle APIs
	http.Handle("/", r)
	// now initializing server here
	log.Fatal(http.ListenAndServe("localhost:4000", r))

}

// So everything is working properly all the 5 API are doing in postman testing
