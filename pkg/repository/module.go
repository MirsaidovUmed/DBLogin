package repository

import "github.com/jackc/pgx/v5"

type Repository struct {
	db *pgx.Conn
}

type RepositoryInterface interface {
	Login(username, password string) error
}

func NewRepository(db *pgx.Conn) RepositoryInterface {
	return &Repository{
		db: db,
	}
}
