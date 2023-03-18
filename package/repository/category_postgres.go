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
		tx.Rollback()
		return 0, err
	}

	createUserCategoryQuery := fmt.Sprintf("INSERT INTO %s (user_id, category_id) VALUES ($1, $2)", usersCategoryTable)
	_, err = tx.Exec(createUserCategoryQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *CategoryProductPostgres) GetAll(userId int) ([]endpoint.Category, error) {
	var categories []endpoint.Category

	query := fmt.Sprintf("SELECT tc.id, tc.title, tc.description FROM %s tc INNER JOIN %s uc on tc.id = uc.category_id where uc.user_id = $1", categoryTable, usersCategoryTable)
	err := r.db.Select(&categories, query, userId)

	return categories, err
}

func (r *CategoryProductPostgres) GetById(userId, categoryId int) (endpoint.Category, error) {
	var category endpoint.Category

	query := fmt.Sprintf(`SELECT tc.id, tc.title, tc.description FROM %s tc INNER JOIN %s uc on tc.id = uc.category_id where uc.user_id = $1 AND uc.category_id = $2`, categoryTable, usersCategoryTable)
	err := r.db.Get(&category, query, userId, categoryId)

	return category, err
}
