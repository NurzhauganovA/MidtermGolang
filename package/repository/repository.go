package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Category interface {
}

type Product interface {
}

type Repository struct {
	Authorization
	Category
	Product
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
