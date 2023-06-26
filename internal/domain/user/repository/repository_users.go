package repository_user

import (
	"database/sql"
	"errors"

	"github.com/eneassena10/estoque-go/internal/domain/user/entities"
)

type Repository struct {
	Database *sql.DB
}

func NewRepository(database *sql.DB) entities.IRepositoryUser {
	r := new(Repository)
	r.Database = database
	return r
}

func (r *Repository) Logar(user entities.User) (bool, error) {
	query := "SELECT nickname, password FROM users WHERE nickname=?"
	stmt := r.Database.QueryRow(query, &user.Nickname)

	var userResult entities.User
	err := stmt.Scan(&userResult.Nickname, &userResult.Password)
	if err != nil {
		return false, errors.New("credenciais ivalida")
	}

	if userResult.Nickname == user.Nickname && userResult.Password == user.Password {
		return true, nil
	}
	return false, errors.New("credenciais ivalida")
}

func (r *Repository) Create(user entities.User) (entities.LoginRequest, error) {
	query := "INSERT INTO users (name, nickname, password,logado) VALUES(?, ?, ?, ?)"
	stmt, err := r.Database.Prepare(query)
	if err != nil {
		return entities.LoginRequest{}, errors.New("dados incosistente")
	}
	defer stmt.Close()

	result, err := stmt.Exec(&user.Name, &user.Nickname, &user.Password, &user.Logado)
	if err != nil {
		return entities.LoginRequest{}, errors.New("dados incosistente")
	}
	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 && err != nil {
		return entities.LoginRequest{}, errors.New("dados incosistente")
	}

	return entities.LoginRequest{
		Name:     user.Name,
		Nickname: user.Nickname,
		Password: user.Password,
	}, nil
}

func (r *Repository) Logout(user entities.User) error {
	return nil
}
