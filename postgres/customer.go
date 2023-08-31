package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/agpelkey/order-service/domain"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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
        VALUES (@user_name, @email, @password)
        RETURNING id
    `
    
    args := pgx.NamedArgs{
        "user_name": &user.Username,
        "email":    &user.Email,
        "password": &user.Password,
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
func (u customerStore) GetAllUsers(ctx context.Context) ([]domain.Customer, error) {
    query := `
        SELECT id, username, email, password FROM customers
    `

    rows, err := u.db.Query(ctx, query)    
    if err != nil {
       return nil, err 
    }

    users, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Customer])
    if err != nil {
        return nil, err
    }

    if len(users) == 0 {
        return nil, domain.ErrNoUsersFound
    }

    return users, nil
}

// Get user by ID
func (u customerStore) GetCustomerByID(ctx context.Context, id int64) (domain.Customer, error) {
    query := `
        SELECT username, email, password FROM customers WHERE id = $1
    `

    var customer domain.Customer

    err := u.db.QueryRow(ctx, query, id).Scan(
        &customer.Username,
        &customer.Email,
        &customer.Password,
    )

    if err != nil {
        return domain.Customer{}, err
    }

    return customer, nil
}

// Delete user
func (u customerStore) DeleteCustomer(ctx context.Context, id int64) error {
    query := `
        DELETE FROM customers WHERE id = $1
    `

    payload, err := u.db.Exec(ctx, query, id)
    if err != nil {
        return fmt.Errorf("failed to delete from users: %v", err)
    }

    if rows := payload.RowsAffected(); rows != 1 {
        return domain.ErrNoUsersFound
    }

    return nil
}







