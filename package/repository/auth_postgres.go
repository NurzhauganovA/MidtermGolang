package repository

import (
	"fmt"
	"github.com/NurzhauganovA/online-store/endpoint"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user endpoint.User) (int, error) {
	var id int
	var query = fmt.Sprintf("INSERT INTO %s (firstname, surname, username, email, password_hash) values ($1, $2, $3, $4, $5) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Surname, user.Username, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
