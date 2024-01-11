package repository

import "context"

func (repo *Repository) DeactivateResetPassword(reset_password_id int) (err error) {
	_, err = repo.db.Exec(context.Background(), `
		UPDATE user_reset_password
		SET is_active = false
		WHERE id = $1
	`, reset_password_id)

	return
}
