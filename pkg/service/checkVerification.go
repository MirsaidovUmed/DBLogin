package service

import "fmt"

func (s *Service) CheckVerification() {
	var username string

	fmt.Println("Введите имя пользователя")

	fmt.Scan(&username)

	user_id, err := s.repo.GetUserIdByName(username)

	if err != nil {
		fmt.Println(err)
		return
	}

	var check string

	fmt.Println("Это точно вы?")

	fmt.Scan(&check)

	if check == "да" {
		// проверка существует ли такой code
		code, err := s.repo.ShowOTPByUserID(user_id)

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Вот держи свой код, и не забывай снова!!!!!", code)
	}
}
