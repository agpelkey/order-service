package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

// custom type to be used to communicate messages back to the client
type envelope map[string]interface{} 

// set const of 1M to be max size for http request
const maxBytesBodyRead = 1_048_576
// function to read JSON from client request
func fromJSON(w http.ResponseWriter, r *http.Request, v any) error {
    r.Body = http.MaxBytesReader(w, r.Body, maxBytesBodyRead)

    // decode the request into the 'v' variable
    err := json.NewDecoder(r.Body).Decode(v)
    if err != nil {
        var MaxBytesError *http.MaxBytesError
        switch {
        case errors.As(err, &MaxBytesError):
            return errors.New("exceeded 1M request body size")
        default:
            return errors.New("failed to parse request body")
        }
    }

    return nil
}

// function to write json back to the client
func writeJSON(w http.ResponseWriter, v any, status int) error {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    return json.NewEncoder(w).Encode(v)
}
