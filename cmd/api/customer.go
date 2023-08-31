package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/agpelkey/order-service/domain"
	"github.com/julienschmidt/httprouter"
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

// method to get customer by id
func (app *application) handleGetCustomerByID(w http.ResponseWriter, r *http.Request) {
   // get id from url 
   id, err := readIdParam(r)
   if err != nil {
        app.serverErrorResponse(w, r, err)
        return
   }

   // query the db
   customer, err := app.CustomerStore.GetCustomerByID(r.Context(), id)
   if err != nil {
        switch {
        case errors.Is(err, domain.ErrNoUsersFound):
            app.notFoundResponse(w, r)
        default:
            app.serverErrorResponse(w, r, err)
        }
        return
   }

   err = writeJSON(w, http.StatusOK, envelope{"customer:":customer}, nil)
}

// method to get all customer
func (app *application) handleGetAllCustomers(w http.ResponseWriter, r *http.Request) {
    customers, err := app.CustomerStore.GetAllUsers(r.Context())
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    _ = writeJSON(w, http.StatusOK, envelope{"customers":customers}, nil)
}


// method to delete a customer
func (app *application) handleDeleteCustomer(w http.ResponseWriter, r *http.Request) {
    // get id 
    id, err := readIdParam(r)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    err = app.CustomerStore.DeleteCustomer(r.Context(), id)
    if err != nil {
        switch {
        case errors.Is(err, domain.ErrRecordNotFound):
            app.notFoundResponse(w, r)
        default:
            app.serverErrorResponse(w, r, err)
        }
    }

    err = writeJSON(w, http.StatusOK, envelope{"message": "user successfully delted"}, nil)
    if err != nil {
        app.serverErrorResponse(w, r, err)
    }
}


// function to get id from url
func readIdParam(r *http.Request) (int64, error) {
    params := httprouter.ParamsFromContext(r.Context())

    id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
    if err != nil {
        return 0, errors.New("invalid id parameter")
    }

    return id, nil
}






