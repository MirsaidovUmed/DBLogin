package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) Login(username, password string) error {
	var db_password string

	row := r.db.QueryRow(context.Background(), "SELECT password FROM users WHERE username = $1", username)

	err := row.Scan(&db_password)

	if err != nil {
		if err == pgx.ErrNoRows {
			return errors.New("пользователь не найден")
		}

		return err
	}

	if db_password != password {
		return errors.New("неправильный пароль")
	}

	return nil
}

// db.Query // получить несколько записей из бд
// db.QueryRow // получить одну запись из бд
//db.Exec // отправить запрос но не ожидать ничего из бд
