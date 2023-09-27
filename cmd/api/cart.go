package main

import (
	"errors"
	"fmt"
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
func (app *application) handleGetCartById(w http.ResponseWriter, r *http.Request) {
    id, err := readIdParam(r)
    if err != nil {
        app.ErrorInvalidQuery(w, r)
        return
    }

    cart, err := app.CartStore.GetCartByID(r.Context(), id)
    if err != nil {
        app.notFoundResponse(w, r)
        fmt.Println(err)
        return
    }

    _ = writeJSON(w, http.StatusOK, envelope{"cart":cart}, nil)
}

// handle update cart
func (app *application) handleUpdateCart(w http.ResponseWriter, r *http.Request) {
    id, err := readIdParam(r)
    if err != nil {
        app.ErrorInvalidQuery(w, r)
        return
    }

    input := domain.CartUpdate{}
    err = readJSON(w, r, &input)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    err = input.Validate()
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    err = app.CartStore.UpdateCart(r.Context(), id, input)
    if err != nil {
        switch {
        case errors.Is(err, domain.ErrNoCartsFound):
            app.notFoundResponse(w, r)
            return
        default:
            app.serverErrorResponse(w, r, err)
            return
        }
    }

    _ = writeJSON(w, http.StatusOK, envelope{"Update":"Cart updated"}, nil)
}


// handle delete cart
func (app *application) handleDeleteCart(w http.ResponseWriter, r *http.Request) {
    id, err := readIdParam(r)
    if err != nil {
        app.ErrorInvalidQuery(w, r)
        return
    }

    cart := app.CartStore.DeleteCart(r.Context(), id)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    _ = writeJSON(w, http.StatusOK, envelope{"cart deleted": cart}, nil)
}















