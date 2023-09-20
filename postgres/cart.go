package postgres

import (
	"context"

	"github.com/agpelkey/order-service/domain"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)


type cartStore struct {
    db *pgxpool.Pool
}

func NewCartStore(db *pgxpool.Pool) cartStore {
    return cartStore{db: db}
}


// create
func (c cartStore) CreateNewCart(ctx context.Context, input *domain.Cart) error {
    query := `
        INSERT INTO carts (entree_id, quantity, customer_id)
        VALUES (@entree_id, @quantity, @customer_id) RETURNING id
    `

    args := pgx.NamedArgs{
        "entree_id": input.EntreeID,
        "quantity": input.Quantity,
        "customer_id": input.CustomerID,
    }

    err := c.db.QueryRow(ctx, query, args).Scan(&input.ID)
    if err != nil {
        pgErr := pgError(err)
        if pgErr.Code == pgerrcode.ForeignKeyViolation {
            if pgErr.ConstraintName == "carts_customer_id_fkey" {
                return domain.ErrCartInvalidCustomerID
            }
            if pgErr.ConstraintName == "carts_entree_id_pkey" {
                return domain.ErrCartInvalidEntreeID
            }
        }

        return err
    }

    return nil
}

// get by id

// get all carts

// update cart

// delete cart
