package domain

import (
	"context"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var (

    ErrDuplicateCustomerEmail = errors.New("duplicated email")
    ErrNoUsersFound = errors.New("no users to list")
    ErrRecordNotFound = errors.New("record not found")

	errEmailRequired = errors.New("email is required")
	errEmailTooLong = errors.New("email is too long")
	errEmailInvalid = errors.New("email is invalid")
	errPasswordRequired = errors.New("password is required")
	errPasswordTooLong = errors.New("password is too long")
	errPasswordTooShort = errors.New("password is too short")
) 

// declare customer type
type Customer struct {
	ID 	 int `json:"id"`
	Username string `json:"user_name"`
	Email 	 string `json:"email"`	
	Password []byte`json:"password"`
}

// declare customer create 
type CustomerCreate struct {
	Username string `json:"user_name"`
	Email  	 string `json:"email"`
	Password string `json:"password"`
} 

// declare customer login 
type CustomerLogin struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type CustomerService interface {
	// put db method logic here
    CreateNewUser(user *Customer) error
    GetAllUsers(ctx context.Context) ([]Customer, error)
    GetCustomerByID(ctx context.Context, id int64) (Customer, error)
    DeleteCustomer(ctx context.Context, id int64) error
}

// Validate to validate POST requests 
func (c CustomerCreate) Validate() error {
	switch {
	case c.Email == "":
		return errEmailRequired
	case len(c.Email) >= 500:
		return errEmailTooLong
	case !strings.Contains(c.Email, "@"):
		return errEmailInvalid
	case c.Password == "":
		return errPasswordRequired
	case len(c.Password) >= 72:
		return errPasswordTooLong
	case len(c.Password) <= 8:
		return errPasswordTooShort
	}

	return nil
}

// Validate PATCH requests model
func (c CustomerLogin) Validate() error {
	switch {
	case c.Email == "":
		return errEmailRequired
	case c.Password == "":
		return errPasswordRequired
	}

	return nil
}

// function to set input values and create new struct 
func (c CustomerCreate) CreateModel(password []byte) Customer {
	return Customer{
		Username: c.Username,
		Email: c.Email,
		Password: password,
	}
}

// Generate a hashed password for the customer
func GenerateHashedPassword(password []byte) ([]byte, error) {
    return bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
}

// Compare hashed and plain text passowrds
func CompareHashAndPassword(hashedPassword []byte, password []byte) error {
    return bcrypt.CompareHashAndPassword(hashedPassword, password)
}




