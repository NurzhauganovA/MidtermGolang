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
	GetAll(userId int) ([]endpoint.Category, error)
	GetById(userId, categoryId int) (endpoint.Category, error)
}

type Product interface {
	Create(categoryId int, product endpoint.Product) (int, error)
	GetAll(userId, categoryId int) ([]endpoint.Product, error)
	GetById(userId, productId int) (endpoint.Product, error)
	GetQueryParam(userId, productTitle string) (endpoint.Product, error)
	Rate(userId, categoryId int) (endpoint.Product, error)
	GetFilteredProducts(price float64, price2 float64, rating int) (interface{}, interface{})
	CreateComment(comment int, id int) []endpoint.Product
}

type Repository struct {
	Authorization
	Category
	Product
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Category:      NewCategoryProductPostgres(db),
		Product:       NewProductPostgres(db),
	}
}
