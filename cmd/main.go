package main

import (
	"fmt"
	"sql_go/pkg/config"
	"sql_go/pkg/database"
	"sql_go/pkg/repository"
	"sql_go/pkg/service"
)

func main() {

	config := config.NewConfig()

	db := database.NewDatabase(config.DatabaseURL)

	repo := repository.NewRepository(db)

	svc := service.NewService(repo)

	for {
		var choice int
		fmt.Println("1. Логин")
		fmt.Println("2. Выйти")

		fmt.Scan(&choice)

		if choice == 1 {
			svc.Login()
		} else if choice == 2 {
			break
		}
	}
}
