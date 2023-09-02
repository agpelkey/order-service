package postgres

import (
	"context"
	"time"

	"github.com/agpelkey/order-service/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type entreeStore struct {
    db *pgxpool.Pool
}

func NewEntreeStore(db *pgxpool.Pool) entreeStore {
    return entreeStore{db: db}
}

// method to create new entree
func (e entreeStore) CreateEntree(entree *domain.Entree) error {
    query := `
        INSERT INTO entrees (name, description, quantity, cost) VALUES (@name, @description, @quantity, @cost)
    `
 
    args := pgx.NamedArgs {
        "name": &entree.Name,
        "description": &entree.Description,
        "quantity": &entree.Quantity,
        "cost": &entree.Cost,
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    err := e.db.QueryRow(ctx, query, args).Scan(&entree.ID)
    if err != nil {
        return err
    }

    return nil
}

// Get by ID

// Get all

// Update

// Delete
