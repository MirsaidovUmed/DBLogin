package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (repo *Repository) GetUserIdByName(username string) (id int, err error) {

	row := repo.db.QueryRow(context.Background(), "SELECT id from users where username = $1", username)

	err = row.Scan(&id)

	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, errors.New("ПОЛЬЗОВАТЕЛЬ НЕ НАЙДЕН!! ")
		}

		return 0, err
	}

	return
}
