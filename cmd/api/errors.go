package main

import (
	"fmt"
	"log"
	"net/http"
)

// generic helper for logging an error message
func (app *application) logError(r *http.Request, err error) {
    log.Printf("[ERROR]: %s %s: %s", r.Method, r.URL, err) 
}

// errorResponse method is a generic helper for sending JSON-formatted
// error messages to the client
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
    env := envelope{"error": message}

    err := writeJSON(w, http.StatusInternalServerError, env, nil)
    if err != nil {
        app.logError(r, err)
        w.WriteHeader(500)
    }
}

// serverErrorResponse method will be used when the application encounters an error at run-time.
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
    app.logError(r, err)

    message := "the server encountered a problem and could not process your request"
    app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// notFoundResponse method will be used to send a 404 Not Found status code
// and JSON response to the client
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
    message := "the requested resource could not be found"
    app.errorResponse(w, r, http.StatusNotFound, message)
} 

// methodNotAllowed method will be used to send a 405 Method Not Allowed status code
// and JSON response to the client
func (app *application) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
    message := fmt.Sprintf("the %s method is not allowed", r.Method) 
    app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

// ErrorInvalidQuery uses Error to report invalid URL queries
func (app *application) ErrorInvalidQuery(w http.ResponseWriter, r *http.Request) {
    app.errorResponse(w, r, http.StatusBadRequest, "invalid url query")
}
