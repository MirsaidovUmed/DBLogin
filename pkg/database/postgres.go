package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func NewDatabase(url string) *pgx.Conn {
	db, err := pgx.Connect(context.Background(), url)

	if err != nil {
		panic(err)
	}

	return db
}
