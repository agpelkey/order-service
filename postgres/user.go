package postgres

import (
	"github.com/agpelkey/order-service/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type customerStore struct {
    db *pgxpool.Pool
}

func NewCustomerStore(db *pgxpool.Pool) customerStore {
    return customerStore{db: db}
}


// Get all users
func (u customerStore) GetAllUsers() ([]domain.Customer, error) {
    query := `
        SELECT * FROM users
    `
}

// Get user by ID

// Create new user

// Update user

// Delete user
