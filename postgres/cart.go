package postgres

import "github.com/jackc/pgx/v5/pgxpool"


type cartStore struct {
    db *pgxpool.Pool
}

func NewCartStore(db *pgxpool.Pool) cartStore {
    return cartStore{db: db}
}
