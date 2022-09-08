package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	code int
	str  string
}

var ErrorWrongMethod = Error{400, "wrong method"}
var ErrorCanNotParseDate = Error{400, "can't parse date"}
var ErrorInternalError = Error{500, "internal error"}

func throwError(w http.ResponseWriter, error Error) {
	w.WriteHeader(error.code)
	v, err := json.Marshal(APIError{Error: error.str})
	if err != nil {
		return
	}
	fmt.Fprint(w, string(v))
}