package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/agpelkey/order-service/domain"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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
func (c cartStore) GetCartByID(ctx context.Context, id int64) (domain.Cart, error) {
    query := `
        SELECT entree_id, quantity, customer_id FROM carts WHERE id = @id
    `

    var cart domain.Cart

    err := c.db.QueryRow(ctx, query).Scan(
        &cart.EntreeID,
        &cart.Quantity,
        &cart.CustomerID,
    )
    if err != nil {
        var pgErr *pgconn.PgError
        switch {
        case errors.Is(err, domain.ErrNoCartsFound):
            return domain.Cart{}, fmt.Errorf(pgErr.Message)
        default:
            return domain.Cart{}, err
        }
    }

    return cart, nil
}

// get all carts

// update cart
func (c cartStore) UpdateCart(ctx context.Context, id int64, input domain.CartUpdate) error {
    query := `
        UPDATE carts 
        SET entree_id = COALESCE(@product_id, product_id),
            quantity = COALESCE(@quantity, quantity),
            customer_id = COALESCE(@customer_id, customer_id),
    `
    
    args := pgx.NamedArgs{
        "entree_id": &input.EntreeID,
        "quantity": &input.Quantity,
        "id": &id,
    }

    _, err := c.db.Query(ctx, query, args)
    if err != nil {
        fmt.Errorf("failed to update cart: %v", err)
    }

    return nil
}

// delete cart
