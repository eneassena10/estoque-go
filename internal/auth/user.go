package auth

import (
	"net/http"

	"github.com/eneassena10/estoque-go/internal/domain"
	"github.com/gin-gonic/gin"
)

type (
	UserController  struct{}
	IUserController interface {
		Logar(ctx *gin.Context)
		Logout(ctx *gin.Context)
		Create(ctx *gin.Context)
	}
	Response struct {
		Data  interface{}
		Error string
	}
)

func NewUserController() IUserController {
	return &UserController{}
}

func (uc *UserController) Logar(ctx *gin.Context) {
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
