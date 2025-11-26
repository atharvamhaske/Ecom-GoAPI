package json

//this file is for encoding and decoding json we are making one global file and will use this

import (
	"net/http"
	"encoding/json"
)

func Write(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func Read(r *http.Request, data any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	
	return decoder.Decode(data)
}