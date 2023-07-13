package controllers

import (
	"github.com/eneassena10/estoque-go/internal/domain/user/entities"
	service_user "github.com/eneassena10/estoque-go/internal/domain/user/service"
	"github.com/eneassena10/estoque-go/pkg/regras"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service *service_user.ServiceUser
}

func NewUserController(service *service_user.ServiceUser) *UserController {
	return &UserController{
		Service: service,
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
