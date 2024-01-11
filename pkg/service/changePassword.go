package service

import "fmt"

func (s *Service) ChangePassword() {
	var username string

	fmt.Println("Введите имя пользователя")

	fmt.Scan(&username)

	// проверка есть ли такой username
	id, err := s.repo.GetUserIdByName(username)

	if err != nil {
		fmt.Println(err)
		return
	}

	var code string

	fmt.Println("введите код который был получен от 544")
	fmt.Scan(&code)

	// проверка существует ли такой code
	reset_password_id, err := s.repo.CheckOTPByUserID(code, id)

	if err != nil {
		fmt.Println(err)
		return
	}

	var new_password string

	fmt.Println("введите новый пароль")
	fmt.Scan(&new_password)

	// изменение пароля пользователя
	err = s.repo.ChangePassword(new_password, id)

	if err != nil {
		fmt.Println(err)
		return
	}

	// не дать ему использовать тот же код
	err = s.repo.DeactivateResetPassword(reset_password_id)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Пароль успешно изменен, больше не забывай, а то мурдем ")

}
