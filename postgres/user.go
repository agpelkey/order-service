package postgres

import "github.com/jackc/pgx/v5/pgxpool"

type userStore struct {
    db *pgxpool.Pool
}

func NewUserStore(db *pgxpool.Pool) userStore {
    return userStore{db: db}
}
