package entities

import "github.com/gin-gonic/gin"

type IUserController interface {
	Logar(ctx *gin.Context)
	Logout(ctx *gin.Context)
	Create(ctx *gin.Context)
}
