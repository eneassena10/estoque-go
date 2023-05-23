package controllers

import (
	"net/http"

	"github.com/eneassena10/estoque-go/internal/user/domain"
	"github.com/gin-gonic/gin"
)

type (
	UserController struct {
		Service domain.IServiceUser
	}
	IUserController interface {
		Logar(ctx *gin.Context)
		Logout(ctx *gin.Context)
		Create(ctx *gin.Context)
	}
)

func NewUserController(service domain.IServiceUser) IUserController {
	return &UserController{
		Service: service,
	}
}

func (uc *UserController) Logar(ctx *gin.Context) {
	var u domain.LoginRequest
	if err := ctx.BindJSON(&u); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Response{Data: u, Error: err.Error()})
		return
	}

	if err := uc.Service.CheckLogin(ctx, u); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Response{Data: u, Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"messager": "login realizado"})
}

func (uc *UserController) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "logout realizado"})
}

func (uc *UserController) Create(ctx *gin.Context) {
	var u domain.User
	if err := ctx.BindJSON(&u); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Response{Data: u, Error: err.Error()})
		return
	}
	u.ID = 1
	ctx.JSON(http.StatusCreated, u)
}
