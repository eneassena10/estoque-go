package user

import (
	repositoryUser "github.com/eneassena10/estoque-go/internal/domain/user/repository"
	serviceUser "github.com/eneassena10/estoque-go/internal/domain/user/service"
	dbsqlite3 "github.com/eneassena10/estoque-go/pkg/conexao/db_sqlite3"
)

type serviceDomoainUser struct {
	Service *serviceUser.ServiceUser
}

func StartDomain(database dbsqlite3.IDataBaseOperation) *serviceDomoainUser {
	repositoryUser := repositoryUser.NewRepository(database)
	serviceUser := serviceUser.NewServiceUser(repositoryUser)
	return &serviceDomoainUser{
		Service: serviceUser,
	}
}
