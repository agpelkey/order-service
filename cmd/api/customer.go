package main

import (
	"fmt"
	"net/http"

	"github.com/agpelkey/order-service/domain"
)

// method to create a new customer
func (app *application) handleCreateCustomer(w http.ResponseWriter, r *http.Request) {
    // create variable to host customer data type
    input := domain.CustomerCreate{}

    // read json into customer create variable
    err := readJSON(w, r, &input)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    // validate that the POST request is not missing any fields
    err = input.Validate()
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    // create a hashed password for the customer
    hashedPassword, err := domain.GenerateHashedPassword([]byte(input.Password))
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    customer := input.CreateModel(hashedPassword)

    // send query to create user in DB
    err = app.CustomerStore.CreateNewUser(&customer)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    headers := make(http.Header)
    headers.Set("Location", fmt.Sprintf("/v1/customers/%d", customer.ID))

    err = writeJSON(w, http.StatusOK, envelope{"customer:": customer}, headers)


}

// method to get all customer

// method to update a customer

// method to delete a customer
