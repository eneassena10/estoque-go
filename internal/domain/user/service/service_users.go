package serviceUser

import (
	"net/http"

	"github.com/eneassena10/estoque-go/internal/domain/user/entities"
	repositoryUser "github.com/eneassena10/estoque-go/internal/domain/user/repository"
	"github.com/eneassena10/estoque-go/pkg/web"

	"github.com/gin-gonic/gin"
)

type ServiceUser struct {
	Repository *repositoryUser.Repository
}

func NewServiceUser(repository *repositoryUser.Repository) *ServiceUser {
	service := ServiceUser{Repository: repository}
	return &service
}

func (s *ServiceUser) Logar(ctx *gin.Context, user entities.LoginRequest) {
	u := entities.NewUser().
		WithName(user.Name).
		WithNickname(user.Nickname).
		WithPassword(user.Password).
		WithLogado(0)

	err := s.Repository.Logar(*u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			web.DecodeError(http.StatusInternalServerError, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusNoContent, nil))
}

func (s *ServiceUser) Logout(ctx *gin.Context, user entities.LoginRequest) {
	err := s.Repository.Logout(entities.User{
		Name:     user.Name,
		Nickname: user.Name,
		Password: user.Password,
		Logado:   1,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			web.DecodeError(http.StatusInternalServerError, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, "OK"))
}

func (s *ServiceUser) Create(ctx *gin.Context, user entities.LoginRequest) {
	userCreate := entities.NewUser().
		WithName(user.Name).
		WithNickname(user.Nickname).
		WithPassword(user.Password).
		WithLogado(0)

	err := s.Repository.Create(*userCreate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			web.DecodeError(http.StatusInternalServerError, err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, nil))
}
