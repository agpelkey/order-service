package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/agpelkey/order-service/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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

// Method to get entree by ID
func (e entreeStore) GetEntreeByID(ctx context.Context, id int64) (domain.Entree, error) {
    query := `
        SELECT name, description, cost, quantity FROM entrees WHERE id = $1
    ` 

    var entree domain.Entree

    err := e.db.QueryRow(ctx, query, id).Scan(
        &entree.Name,
        &entree.Description,
        &entree.Cost,
        &entree.Quantity,
    )

    if err != nil {
        var pgErr *pgconn.PgError
        switch {
        case errors.Is(err, domain.ErrNoEntreesFound):
            return domain.Entree{}, fmt.Errorf(pgErr.Message)
        default:
            return domain.Entree{}, err
        }
    }

    return entree, nil
}

// Method to update entree
func (e entreeStore) UpdateEntreeByID(ctx context.Context, id int64, input domain.Entree) (domain.Entree, error) {

    query := `
       UPDATE entrees 
       SET name         = COALESCE(@name, name),
            description = COALESCE(@description, description),
            cost        = COALESCE(@cost, cost),
            quantity    = COALESCE(@quantity, quantity),
        WHERE id = @id RETURNING *
        `
    
    args := pgx.NamedArgs{
        "name": &input.Name,
        "description": &input.Description,
        "cost": &input.Cost,
        "quantity": &input.Quantity,
    }


    row, err := e.db.Query(ctx, query, args)
    if err != nil {
        return domain.Entree{}, fmt.Errorf("failed to update produce: %v", err)
    }


}


// Delete














