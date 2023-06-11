package controllers

import (
	"database/sql"

	"github.com/eneassena10/estoque-go/internal/domain/user/entities"
	repository_user "github.com/eneassena10/estoque-go/internal/domain/user/repository"
	service_user "github.com/eneassena10/estoque-go/internal/domain/user/service"
	"github.com/eneassena10/estoque-go/pkg/regras"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service entities.IServiceUser
}

func NewUserController(database *sql.DB) entities.IUserController {
	repositoryUser := repository_user.NewRepository(database)
	serviceUser := service_user.NewServiceUser(repositoryUser)
	return &UserController{
		Service: serviceUser,
	}
}

func (uc *UserController) Logar(ctx *gin.Context) {
	var data entities.LoginRequest
	if regras.ValidateErrorInRequest(ctx, &data) {
		return
	}

	uc.Service.Logar(ctx, data)
}

func (uc *UserController) Logout(ctx *gin.Context) {
	var data entities.LoginRequest
	if regras.ValidateErrorInRequest(ctx, &data) {
		return
	}

	uc.Service.Logout(ctx, data)
}

func (uc *UserController) Create(ctx *gin.Context) {
	var data entities.LoginRequest
	if regras.ValidateErrorInRequest(ctx, &data) {
		return
	}

	uc.Service.Create(ctx, data)
}
