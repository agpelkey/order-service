package postgres

import (
	"context"
	"time"

	"github.com/agpelkey/order-service/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

// tokenStore represents token database.
type tokenStore struct {
    db *pgxpool.Pool
}

// NewTokenStore returns a new instance of TokenStore
func NewTokenStore(db *pgxpool.Pool) tokenStore {
    return tokenStore{db: db}
}

// method to create new token
func (t tokenStore) CreateToken(ctx context.Context, customerID int64, ttl time.Duration, scope string) (*domain.Token, error) {

    token, err := domain.GenerateToken(customerID, ttl, scope)
    if err != nil {
        return nil, err
    }

    err = t.Insert(token)
    return token, err
}


// method to add the data for a specific token to the tokens table
func (t tokenStore) Insert(token *domain.Token) error {
    query := `
        INSERT INTO tokens (hash, customer_id, expiry, scope)
        VALUE ($1, $2, $3, $4)
    `

    args := []interface{}{token.Hash, token.CustomerID, token.Expiry, token.Scope}

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := t.db.Exec(ctx, query, args...)
    return err
}

