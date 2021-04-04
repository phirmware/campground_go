package utils

import (
	"errors"
	"net/http"

	"github.com/gorilla/schema"
)

func GetRequestBody(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	dec := schema.NewDecoder()
	if err := dec.Decode(dst, r.PostForm); err != nil {
		return errors.New("Error parsing the request body")
	}
	return nil
}
