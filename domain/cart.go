package domain

import (
	"context"
	"errors"

)

var (
	ErrCartInvalidCustomerID = errors.New("invalid customer id")
	ErrCartInvalidEntreeID   = errors.New("invalid entree id")
	ErrNoCartsFound      = errors.New("no cart could be found")

	errUserIDRequired   = errors.New("customer_id is required")
	errEntreeIDRequired = errors.New("entre_id is required")

    
)

// Cart type
type Cart struct {
	ID       int `json:"id"`
	EntreeID int `json:"entree_id"`
	CustomerID int `json:"customer_id"`
	Quantity int `json:"quantity"`
}

// type for creating new cart
type CartCreate struct {
	ID         int `json:"id"`
	EntreeID   int `json:"entree_id"`
	CustomerID int `json:"customer_id"`
	Quantity   int `json:"quantity"`
}

// type for updating a cart
type CartUpdate struct {
	EntreeID *int `json:"product_id"`
	Quantity *int `json:"quantity"`
}

// CartService represents service for managing carts
type CartService interface {
	// Insert DB methods here
	GetCartByID(ctx context.Context, id int64) (Cart, error)
	CreateNewCart(ctx context.Context, input *Cart) error
	UpdateCart(ctx context.Context, id int64, input CartUpdate) error
	DeleteCart(ctx context.Context, id int64) error
}

// Validate Post request to create cart
func (c CartCreate) Validate() error {
	switch {
	case c.EntreeID == 0:
		return errEntreeIDRequired
	case c.CustomerID == 0:
		return errUserIDRequired
	case c.Quantity == 0:
		return errQuantityRequired
	}

	return nil
}

// reateModel sets input values and returns new struct
func (c CartCreate) CreateModel() Cart {
	return Cart{
		EntreeID: c.EntreeID,
		CustomerID:   c.CustomerID,
		Quantity: c.Quantity,
	}
}

// Validate validates PATCH requests to cart
func (c CartUpdate) Validate() error {
	switch {
	case c.EntreeID != nil && *c.EntreeID == 0:
		return errEntreeIDRequired
	case c.Quantity != nil && *c.Quantity == 0:
		return errQuantityRequired
	}

	return nil
}

// UpdateModel checks whether carts input are not nil and sets values
func (c CartUpdate) UpdateModel(cart *Cart) {
	if c.EntreeID != nil {
		cart.EntreeID = *c.EntreeID
	}

	if c.Quantity != nil {
		cart.Quantity = *c.Quantity
	}
}
