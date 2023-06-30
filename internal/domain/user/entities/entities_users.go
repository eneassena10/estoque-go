package entities

import "github.com/gin-gonic/gin"

type IUserLogout interface {
	Logout(u User) error
}
type IUserCreated interface {
	Create(u User) (User, error)
}
type IUserLogger interface {
	Logar(u User) error
}
type IUser interface {
	IUserLogger
	IUserCreated
	IUserLogout
}
type IRepositoryUser interface {
	Logar(user User) error
	Create(user User) error
	Logout(user User) error
}

type IServiceUser interface {
	Logar(ctx *gin.Context, user LoginRequest)
	Create(ctx *gin.Context, user LoginRequest)
	Logout(ctx *gin.Context, user LoginRequest)
}
type User struct {
	ID       int
	Name     string
	Nickname string
	Password string
	Logado   int
}

type LoginRequest struct {
	Name     string `json:"name" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) WithID(id int) *User {
	u.ID = id
	return u
}

func (u *User) WithName(name string) *User {
	u.Name = name
	return u
}

func (u *User) WithNickname(nickname string) *User {
	u.Nickname = nickname
	return u
}

func (u *User) WithPassword(password string) *User {
	u.Password = password
	return u
}

func (u *User) WithLogado(logado int) *User {
	u.Logado = logado
	return u
}
