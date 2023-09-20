package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/agpelkey/order-service/domain"
)

// method to handle creating a new cart
// @Router /v1/cart/  [POST]
func (app *application) handleCreateCart(w http.ResponseWriter, r *http.Request) {
    cart := domain.CartCreate{}

    err := readJSON(w, r, &cart)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    err = cart.Validate()
    if err != nil {
       app.ErrorInvalidQuery(w, r) 
       return
    }

    log.Println("passed validation")

    result := cart.CreateModel()
    err = app.CartStore.CreateNewCart(r.Context(), &result)
    if err != nil {
        if errors.Is(err, domain.ErrCartInvalidCustomerID) ||
            errors.Is(err, domain.ErrCartInvalidEntreeID) {
            app.ErrorInvalidQuery(w, r) 
            } else {
                app.serverErrorResponse(w, r, err) 
                return
            }
        return
    }

    err = writeJSON(w, http.StatusOK, envelope{"Cart created": result}, nil)
}


// handle get cart by id

// handle update cart

// handle delete cart
