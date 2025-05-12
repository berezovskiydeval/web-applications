package repository

import (
	"database/sql"

	"fmt"
	"github.com/berezovskyivalerii/notes-manager/backend/internal/domain"
	db "github.com/berezovskyivalerii/notes-manager/backend/pkg/database"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user domain.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", db.UsersTable)
	var id int
	err := r.db.QueryRow(query, user.Name, user.Username, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (domain.User, error) {
	var user domain.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 AND password_hash = $2", db.UsersTable)
	err := r.db.QueryRow(query, username, password).Scan(&user.Id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
