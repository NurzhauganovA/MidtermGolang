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
	Create(userId int, category endpoint.Category) (int, error)
	GetAll(userId int) ([]endpoint.Category, error)
	GetById(userId, categoryId int) (endpoint.Category, error)
}

type Product interface {
	Create(userId, categoryId int, product endpoint.Product) (int, error)
	GetAll(userId, categoryId int) ([]endpoint.Product, error)
	GetById(userId, productId int) (endpoint.Product, error)
}

type Service struct {
	Authorization
	Category
	Product
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Category:      NewCategoryProductService(repos.Category),
		Product:       NewProductService(repos.Product, repos.Category),
	}
}
