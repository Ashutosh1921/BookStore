package utils

import (
	"encoding/json"
	"io"
	"net/http"

	models "github.com/Ashutosh1921/BookStore/pkg/Models"
)

func ParseBody(r *http.Request, x *models.Book) {
	// here x is expecting pointer to type book which will be the empty struct
	// inside which we will parse data from htttp request into struct then into the database
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

// this function is basically filling data from http response to our interface
// io.readall is converting json data into slices of bytes
// then json unmarshalling is filling that byte data into struct then it goes to data base
