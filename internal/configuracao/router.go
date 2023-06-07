package configuracao

import (
	productControllers "github.com/eneassena10/estoque-go/internal/domain/product/controllers"
	userControllers "github.com/eneassena10/estoque-go/internal/domain/user/controllers"
	"github.com/gin-gonic/gin"
)

const (
	// Products - rotas de listagem e criação de um novo produto
	Products = "/products"

	// ProductsID - usado em rotas que usa o id do produto para uma alteração e leitura
	ProductsList = "/products/list"

	UserLogin  = "user/login"
	UserLogout = "user/logout"
	UserCreate = "user/create"
)

type App struct {
	products productControllers.IProductControllers
	user     userControllers.IUserController
}
type IApp interface {
	InitApp(router *gin.Engine)
}

func NewApp(products productControllers.IProductControllers, user userControllers.IUserController) IApp {
	return &App{products: products, user: user}
}

func (a *App) InitApp(router *gin.Engine) {
	router.GET(Products, a.products.GetProductsByID)
	router.GET(ProductsList, a.products.GetProductsAll)
	router.DELETE(Products, a.products.DeleteProducts)
	router.POST(Products, a.products.CreateProducts)
	router.PATCH(Products, a.products.UpdateProductsCount)

	router.POST(UserLogin, a.user.Logar)
	router.POST(UserLogout, a.user.Logout)
	router.POST(UserCreate, a.user.Create)
}
