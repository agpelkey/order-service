package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)


func (app *application) routes() http.Handler {
    router := httprouter.New()

    //Insert routes here

    return router
}
