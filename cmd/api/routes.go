package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)


func (app *application) routes() http.Handler {
    router := httprouter.New()

    router.HandlerFunc(http.MethodGet, "/v1/health", app.handleHealthCheck)

    router.HandlerFunc(http.MethodPost, "/v1/customers", app.handleCreateCustomer)

    return router
}
