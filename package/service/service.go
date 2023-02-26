package service

import (
	"github.com/NurzhauganovA/online-store/endpoint"
	"github.com/NurzhauganovA/online-store/package/repository"
)

type Authorization interface {
	CreateUser(user endpoint.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Category interface {
}

type Product interface {
}

type Service struct {
	Authorization
	Category
	Product
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
