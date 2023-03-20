package repository

import (
	"fmt"
	"github.com/NurzhauganovA/online-store/endpoint"
	"github.com/jmoiron/sqlx"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) Create(categoryId int, product endpoint.Product) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var productId int
	createProductQuery := fmt.Sprintf("INSERT INTO %s (title, description, cost, created_country, created_company, created_date) values ($1, $2, $3, $4, $5, $6) RETURNING id", productTable)
	row := tx.QueryRow(createProductQuery, product.Title, product.Description, product.Cost, product.CreatedCountry, product.CreatedCompany, product.CreatedDate)
	err = row.Scan(&productId)

	if err != nil {
		return 0, err
	}

	createListProductQuery := fmt.Sprintf("INSERT INTO %s (category_id, product_id) values ($1, $2)", categoryProductTable)
	_, err = tx.Exec(createListProductQuery, categoryId, productId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return productId, tx.Commit()
}

func (r *ProductPostgres) GetAll(userId, categoryId int) ([]endpoint.Product, error) {
	var products []endpoint.Product
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.cost, ti.created_company, ti.created_country, ti.created_date FROM %s ti INNER JOIN %s li on li.product_id = ti.id 
											  INNER JOIN %s ul on ul.category_id = li.category_id WHERE li.category_id = $1 AND ul.user_id = $2`,
		productTable, categoryProductTable, usersCategoryTable)

	if err := r.db.Select(&products, query, categoryId, userId); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductPostgres) GetById(userId, productId int) (endpoint.Product, error) {
	var product endpoint.Product
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.cost, ti.created_company, ti.created_country, ti.created_date FROM %s ti INNER JOIN %s li on li.product_id = ti.id 
											  INNER JOIN %s ul on ul.category_id = li.category_id WHERE ti.id = $1 AND ul.user_id = $2`,
		productTable, categoryProductTable, usersCategoryTable)

	if err := r.db.Get(&product, query, productId, userId); err != nil {
		return product, err
	}

	return product, nil
}

func (r *ProductPostgres) GetQueryParam(userId, productTitle string) (endpoint.Product, error) {
	var product endpoint.Product
	rows := fmt.Sprintf(`SELECT pr.id, pr.title, pr.cost FROM %s pr WHERE pr.title LIKE ? $1`, productTable)

	if err := r.db.Get(&product.Title, rows, productTitle, userId); err != nil {
		return product, err
	}

	return product, nil
}
