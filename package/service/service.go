package service

import "github.com/NurzhauganovA/online-store/package/repository"

type Authorization interface {
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
	return &Service{}
}
