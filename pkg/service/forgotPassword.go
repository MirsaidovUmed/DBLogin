package service

import (
	"fmt"
	"sql_go/pkg/utils"
)

func (s *Service) ForgotPassword() {
	var username string

	fmt.Println("А ГОЛОВУ ТЫ НЕ ЗАБЫЛ???!")
	fmt.Println("введите имя пользователя")

	fmt.Scan(&username)

	id, err := s.repo.GetUserIdByName(username)

	if err != nil {
		fmt.Println(err)
		return
	}

	otp := utils.GenOTP(4)

	err = s.repo.ForgotPassword(id, otp)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Заявка на сброс пароля успешно сохранена, позвоните по номеру 544 чтобы получить код")
}
