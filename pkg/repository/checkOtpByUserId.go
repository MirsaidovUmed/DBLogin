package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (repo *Repository) CheckOTPByUserID(code string, user_id int) (id int, err error) {

	row := repo.db.QueryRow(
		context.Background(),
		`SELECT 
			id 
		FROM 
			user_reset_password 
		WHERE
				user_id = $1 
			AND 
				code = $2
			AND 
				is_active = true
			AND 
				expired_at > now()`, user_id, code)

	err = row.Scan(&id)

	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, errors.New("доступ запрещен")
		}
	}

	return
}
