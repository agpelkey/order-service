package postgres

import (
	"context"
	"time"

	"github.com/agpelkey/order-service/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
    "github.com/jackc/pgerrcode"
)

type customerStore struct {
    db *pgxpool.Pool
}

func NewCustomerStore(db *pgxpool.Pool) customerStore {
    return customerStore{db: db}
}

// Create new user
func (u customerStore) CreateNewUser(user *domain.Customer) error {
    query := `
        INSERT INTO customers (username, email, password)
        VALUES ($1, $2, $3)
        RETURNING id
    `
    
    args := pgx.NamedArgs{
        "username": &user.Username,
        "email":    &user.Email,
        "password": &user.Passowrd,
    }
    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()

    err := u.db.QueryRow(ctx, query, args).Scan(&user.ID)
    if err != nil {
        pgErr := pgError(err)
        if pgErr.Code == pgerrcode.UniqueViolation {
            if pgErr.ConstraintName == "users_email_key" {
                return domain.ErrDuplicateCustomerEmail
            }
        }
        return err
    }
    return nil
}

// Get all users
func (u customerStore) GetAllUsers() ([]domain.Customer, error) {
    return []domain.Customer{}, nil
}

// Get user by ID

// Update user

// Delete user
