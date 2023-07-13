package repositoryUser

import (
	"errors"

	"github.com/eneassena10/estoque-go/internal/domain/user/entities"
	dbsqlite3 "github.com/eneassena10/estoque-go/pkg/conexao/db_sqlite3"
)

type Repository struct {
	database dbsqlite3.IDataBaseOperation
}

func NewRepository(database dbsqlite3.IDataBaseOperation) *Repository {
	return &Repository{database: database}
}

func (r *Repository) Logar(user entities.User) error {
	// query := "SELECT nickname, password FROM users WHERE nickname=?"
	// stmt := r.database.QueryRow(query, &user.Nickname)

	// var userResult entities.User
	// err := stmt.Scan(&userResult.Nickname, &userResult.Password)
	// if err != nil {
	// 	return errors.New("credenciais ivalida")
	// }

	// if userResult.Nickname == user.Nickname && userResult.Password == user.Password {
	// 	return nil
	// }
	return errors.New("credenciais ivalida")
}

func (r *Repository) Create(user entities.User) error {
	// query := "INSERT INTO users (name, nickname, password,logado) VALUES(?, ?, ?, ?)"
	// stmt, err := r.database.Prepare(query)
	// if err != nil {
	// 	return errors.New("dados incosistente")
	// }
	// defer stmt.Close()

	// result, err := stmt.Exec(&user.Name, &user.Nickname, &user.Password, &user.Logado)
	// if err != nil {
	// 	return errors.New("dados incosistente")
	// }
	// if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 && err != nil {
	// 	return errors.New("dados incosistente")
	// }

	return nil
}

func (r *Repository) Logout(user entities.User) error {
	return nil
}
