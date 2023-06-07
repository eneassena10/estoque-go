package service_user

import (
	"database/sql"
	"errors"

	"github.com/eneassena10/estoque-go/internal/domain/user/domain"

	"github.com/gin-gonic/gin"
)

type ServiceUser struct {
	DataBase *sql.DB
}

func NewServiceUser(db *sql.DB) domain.IServiceUser {
	service := ServiceUser{DataBase: db}
	return &service
}

func (s *ServiceUser) CheckLogin(ctx *gin.Context, user domain.LoginRequest) error {
	for i := range domain.UsersRegistred {
		if domain.UsersRegistred[i].Nickname == user.Nickname && domain.UsersRegistred[i].Password == user.Password {
			return nil
		}
	}
	return errors.New("credenciais invalida")
}
