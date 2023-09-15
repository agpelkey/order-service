package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/agpelkey/order-service/domain"
)

// @Summary    Get product by ID
// @Tags       Entrees
// @Produce    JSON
// @Param      id       path      int true "Entree ID"
// @Success    200      {array}   envelope{"entree":entree}
// @Failure    400
// @Failure    404
// @Failure    500
// @Router     /entrees/id [get]
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


// @Summary   Create entree
// @Tags      Entrees
// @Produce   JSON
// @Accept    JSON
// @Param     entree  body   domain.CreateEntree true "Create entree"
// @Success   200
// @Failure   400
// @Failure   500 
// @Router    /entrees [post]
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

// @Summary   Update entree
// @Tags      Entrees
// @Produce   JSON
// @Accept    JSON
// @Param     entree  body   domain.CreateEntree true "Create entree"
// @Success   200
// @Failure   404
// @Failure   500 
// @Router    /entrees/:id [patch]
func (app *application) handleUpdateEntree(w http.ResponseWriter, r *http.Request) {
	id, err := readIdParam(r)
	if err != nil {
		app.serverErrorResponse(w, r, err)
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


	entree, err := app.EntreeStore.UpdateEntreeByID(r.Context(), id, input)
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


	err = writeJSON(w, http.StatusOK, envelope{"entree": entree}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}

// delete entree
















