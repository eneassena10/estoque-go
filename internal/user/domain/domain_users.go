package domain

import "github.com/gin-gonic/gin"

type User struct {
	ID       int
	Name     string
	Nickname string
	Password string
}
type IUser interface {
	Logar(u User) error
	Logout(u User) error
	Create(u User) (User, error)
}

type IServiceUser interface {
	CheckLogin(ctx *gin.Context, user LoginRequest) error
}

type LoginRequest struct {
	Name     string `json:"name" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var ProductPathName = "../../data/users.json"

var UsersRegistred []*User = []*User{
	{
		ID:       1,
		Name:     "user test 1",
		Nickname: "user1",
		Password: "user1",
	},
	{
		ID:       2,
		Name:     "user test 2",
		Nickname: "user2",
		Password: "user2",
	},
}
