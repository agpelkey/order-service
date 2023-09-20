package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/agpelkey/order-service/domain"
)

// method to handle getting an entree by ID 
// @Router /v1/entree/:id  [GET]
func (app *application) handleGetEntreeByID(w http.ResponseWriter, r *http.Request) {
    id, err := readIdParam(r)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    payload, err := app.EntreeStore.GetEntreeByID(r.Context(), id)
    if err != nil {
        switch {
        case errors.Is(err, domain.ErrNoEntreesFound):
            app.notFoundResponse(w, r)
        default:
            app.serverErrorResponse(w, r, err)
        }
    }

    err = writeJSON(w, http.StatusOK, envelope{"entree":payload}, nil)

}


// method to handle creating a new entree 
// @Router /v1/entrees/  [POST]
func (app *application) handleCreateEntree(w http.ResponseWriter, r *http.Request) {
    input := domain.EntreeCreate{}

    err := readJSON(w, r, &input)
    if err != nil {
        app.errorResponse(w, r, http.StatusBadRequest, err)
        return
    }

    err = input.Validate()
    if err != nil {
        app.errorResponse(w, r, http.StatusBadRequest, err)
        return
    }

    newEntree := input.CreateModel()

    err = app.EntreeStore.CreateEntree(&newEntree)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    headers := make(http.Header)
    headers.Set("Location", fmt.Sprintf("/v1/entrees/%d", newEntree.ID))

    err = writeJSON(w, http.StatusOK, envelope{"entree":newEntree}, headers)
}

// get all entrees

// method to handle updating an entree 
// @Router /v1/entrees/:id  [PATCH]
func (app *application) handleUpdateEntree(w http.ResponseWriter, r *http.Request) {
	id, err := readIdParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	input := domain.EntreeUpdate{}
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

    err = app.EntreeStore.UpdateEntreeByID(r.Context(), id, input)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNoEntreesFound):
			app.notFoundResponse(w, r)
			return
		default:
			app.serverErrorResponse(w, r, err)
			return
		}
	}


	err = writeJSON(w, http.StatusOK, envelope{"entree": err}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}


// method to handle deleting an entree 
// @Router /v1/entrees/:id  [DELETE]
func (app *application) handleDeleteEntree(w http.ResponseWriter, r *http.Request) {
    id, err := readIdParam(r)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    err = app.EntreeStore.DeleteEntreeByID(r.Context(), id)
    if err != nil {
        if errors.Is(err, domain.ErrNoEntreesFound) {
            app.notFoundResponse(w, r)
            return
        } else {
            app.serverErrorResponse(w, r, err)
            return
        }
    }

    _ = writeJSON(w, http.StatusOK, envelope{"message": "entree deleted"}, nil)

}















