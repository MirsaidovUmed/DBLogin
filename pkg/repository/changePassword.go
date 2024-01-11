package repository

import "context"

func (repo *Repository) ChangePassword(new_password string, user_id int) (err error) {
	_, err = repo.db.Exec(context.Background(), `
		UPDATE 
			users 
		SET 
			password = $1 
		WHERE 
			id = $2
		`, new_password, user_id)

	return
}
