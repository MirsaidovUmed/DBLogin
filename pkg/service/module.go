package service

import "sql_go/pkg/repository"

type Service struct {
	repo repository.RepositoryInterface
}

type ServiceInterface interface {
	Login()
	ForgotPassword()
	ChangePassword()
	CheckVerification()
}

func NewService(repo repository.RepositoryInterface) ServiceInterface {
	return &Service{
		repo: repo,
	}
}
