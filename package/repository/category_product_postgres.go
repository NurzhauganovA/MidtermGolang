package repository

import (
	"fmt"
	"github.com/NurzhauganovA/online-store/endpoint"
	"github.com/jmoiron/sqlx"
)

type CategoryProductPostgres struct {
	db *sqlx.DB
}

func NewCategoryProductPostgres(db *sqlx.DB) *CategoryProductPostgres {
	return &CategoryProductPostgres{db: db}
}

func (r *CategoryProductPostgres) Create(userId int, category endpoint.Category) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createCategoryQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", categoryTable)
	row := tx.QueryRow(createCategoryQuery, category.Title, category.Description)
	if err := row.Scan(&id); err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	return id, tx.Commit()
}
