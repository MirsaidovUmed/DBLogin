package repository

import "context"

func (repo *Repository) ForgotPassword(user_id int, code string) (err error) {
	_, err = repo.db.Exec(
		context.Background(),
		"UPDATE user_reset_password SET is_active = false WHERE user_id = $1",
		user_id,
	)
	if err != nil {
		return err
	}

	_, err = repo.db.Exec(
		context.Background(),
		"INSERT INTO user_reset_password(code, user_id, expired_at) VALUES ($1, $2, now() + interval '1' day)",
		code,
		user_id,
	)
	if err != nil {
		return err
	}

	return nil
}
