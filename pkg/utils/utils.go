package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func parseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

// this function is basically filling data from http response to our interface
// io.readall is converting json data into slices of bytes
// then json unmarshalling is filling that byte data into struct then it goes to data base
