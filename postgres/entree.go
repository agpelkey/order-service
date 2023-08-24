package postgres

import "github.com/jackc/pgx/v5/pgxpool"

type entreeStore struct {
    db *pgxpool.Pool
}

func NewEntreeStore(db *pgxpool.Pool) entreeStore {
    return entreeStore{db: db}
}
