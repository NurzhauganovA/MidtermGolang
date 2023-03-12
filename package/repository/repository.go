package repository

import (
	"github.com/NurzhauganovA/online-store/endpoint"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user endpoint.User) (int, error)
	GetUser(username, password string) (endpoint.User, error)
}

type Category interface {
	Create(userId int, category endpoint.Category) (int, error)
}

type Product interface {
}

type Repository struct {
	Authorization
	Category
	Product
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
