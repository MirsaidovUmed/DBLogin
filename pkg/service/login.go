package service

import "fmt"

func (s *Service) Login() {
	var username, password string

	fmt.Println("Введите имя пользователя")

	fmt.Scan(&username)

	fmt.Println("Введите пароль")

	fmt.Scan(&password)

	err := s.repo.Login(username, password)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Успешно")
}
