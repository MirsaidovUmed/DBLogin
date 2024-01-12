package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

func (repo *Repository) ShowOTPByUserID(user_id int) (code string, err error) {
	row := repo.db.QueryRow(
		context.Background(),
		`SELECT
				code
			FROM
			    user_reset_password
			WHERE 
			    user_id = $1
			AND
			    is_active = true
			AND 
				expired_at > now()`, user_id)
	err = row.Scan(&code)

	if err != nil {
		if err == pgx.ErrNoRows {
			return code, errors.New("Доступ запрещён! ")
		}
		return code, err
	}
	return
}
