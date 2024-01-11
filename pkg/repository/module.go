package repository

import "github.com/jackc/pgx/v5"

type Repository struct {
	db *pgx.Conn
}

type RepositoryInterface interface {
	Login(username, password string) error
	ChangePassword(new_password string, user_id int) (err error)
	CheckOTPByUserID(code string, user_id int) (id int, err error)
	DeactivateResetPassword(reset_password_id int) (err error)
	ForgotPassword(user_id int, code string) (err error)
	GetUserIdByName(username string) (id int, err error)
	ShowOTPByUserID(user_id int) (code string, err error)
}

func NewRepository(db *pgx.Conn) RepositoryInterface {
	return &Repository{
		db: db,
	}
}
