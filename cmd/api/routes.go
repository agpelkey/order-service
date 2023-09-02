package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)


func (app *application) routes() http.Handler {
    router := httprouter.New()

    // healtch check routes
    router.HandlerFunc(http.MethodGet, "/v1/health", app.handleHealthCheck)

    // customer routes
    router.HandlerFunc(http.MethodPost, "/v1/customers", app.handleCreateCustomer)
    router.HandlerFunc(http.MethodGet, "/v1/customers", app.handleGetAllCustomers)
    router.HandlerFunc(http.MethodGet, "/v1/customers/:id", app.handleGetCustomerByID)
    router.HandlerFunc(http.MethodDelete, "/v1/customers/:id", app.handleDeleteCustomer)

    // entree routes
    router.HandlerFunc(http.MethodPost, "/v1/entrees", app.handleCreateEntree)

    return router
}
