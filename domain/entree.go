package domain

import (
	"context"
	"errors"
)

var (

	ErrDuplicateEntree = errors.New("duplicated product")
	ErrNoEntreesFound = errors.New("products not found")
    ErrNoRecordFound = errors.New("record not found")

	errEntreeNameRequired = errors.New("name is required")
	errEntreeDescriptionRequired = errors.New("description is required")
	errQuantityRequired = errors.New("quantity is required")
	errPriceRequired = errors.New("price is required")
	
)

type Entree struct {
	ID          int `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Cost        int `json:"cost"`
	Quantity    int `json:"quantity"`
}

// model to create a new entree
type EntreeCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Cost        int `json:"cost"`
	Quantity    int `json:"quantity"`
}

// model to update entree
type EntreeUpdate struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Cost        *int `json:"cost"`
	Quantity    *int `json:"quantity"`
}

// repository pattern for entree's
type EntreeService interface {
    CreateEntree(entree *Entree) error
	// Create
    GetEntreeByID(ctx context.Context, id int64) (Entree, error)
	// GetByName
	// Update
	// Delete
}

// Used to validate POST requests
func (e EntreeCreate) Validate() error {
	switch {
	case e.Name == "":
		return errEntreeNameRequired
	case e.Description == "":
		return errEntreeDescriptionRequired
	case e.Cost == 0:
		return errPriceRequired
	case e.Quantity == 0:
		return errQuantityRequired
	}

	return nil
}

// CreateModel takes input values and returns a new struct
func (e EntreeCreate) CreateModel() Entree {
	return Entree{
		Name: e.Name,
		Description: e.Description,
		Cost: e.Cost,
		Quantity: e.Quantity,
	}
}

// Used to validate PATCH requests
func (e EntreeUpdate) Validate() error {
	switch {
	case e.Name != nil && *e.Name == "":
		return errEntreeNameRequired
	case e.Description != nil && *e.Description == "":
		return errEntreeDescriptionRequired
	case e.Cost != nil && *e.Cost == 0:
		return errPriceRequired
	case e.Quantity != nil && *e.Quantity == 0:
		return errQuantityRequired
	}

	return nil
}

// UpdateModel checks whether the products input are not nil and sets the values
func (e EntreeUpdate) UpdateModel(entree *Entree) {
	if e.Name != nil {
		entree.Name = *e.Name	
	}	

	if e.Description != nil {
		entree.Description = *e.Description
	}	

	if e.Cost != nil {
		entree.Cost = *e.Cost
	}	
	
	if e.Quantity != nil {
		entree.Quantity = *e.Quantity
	}	

}























