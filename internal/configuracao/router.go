package configuracao

import (
	productControllers "github.com/eneassena10/estoque-go/internal/domain/product/controllers"
	userControllers "github.com/eneassena10/estoque-go/internal/domain/user/controllers"
	"github.com/eneassena10/estoque-go/pkg/store"
	"github.com/gin-gonic/gin"
)

const (
	// Products - rotas de listagem e criação de um novo produto
	Products = "/products"

	// ProductsID - usado em rotas que usa o id do produto para uma alteração e leitura
	ProductsList = "/products/list"
)

type App struct {
	fileStore store.IStore
	products  productControllers.IControllers
	user      userControllers.IUserController
}
type IApp interface {
	InitApp(router *gin.Engine)
}

func NewApp(fs store.IStore, products productControllers.IControllers, user userControllers.IUserController) IApp {
	return &App{fileStore: fs, products: products, user: user}
}

func (a *App) InitApp(router *gin.Engine) {
	router.GET(Products, a.products.GetProductsByID)
	router.GET(ProductsList, a.products.GetProductsAll)
	router.DELETE(Products, a.products.DeleteProducts)
	router.POST(Products, a.products.CreateProducts)
	router.PATCH(Products, a.products.UpdateProductsCount)

	router.POST("user/login", a.user.Logar)
	router.POST("user/logout", a.user.Logout)
	router.POST("user/create", a.user.Create)
}
