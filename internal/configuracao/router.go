package configuracao

import (
	productControllers "github.com/eneassena10/estoque-go/internal/domain/product/controllers"
	"github.com/gin-gonic/gin"
)

var (
	// Products - rotas de listagem e criação de um novo produto
	Products = "/products"

	// ProductsID - usado em rotas que usa o id do produto para uma alteração e leitura
	ProductsList = "/products/list"

	UserLogin  = "/user/login"
	UserLogout = "/user/logout"
	UserCreate = "/user/create"
)

type App struct {
	products *productControllers.ProductControllers
}

func NewApp(products *productControllers.ProductControllers) *App {
	app := &App{products: products}

	return app
}

func (a *App) InitApp(router *gin.Engine) {
	router.GET(ProductsList, a.products.GetProductsAll)
	router.GET(Products, a.products.GetProductsByID)
	router.DELETE(Products, a.products.DeleteProducts)
	router.POST(Products, a.products.CreateProducts)
	router.PATCH(Products, a.products.UpdateProductsCount)

	// router.POST(UserLogin, a.user.Logar)
	// router.POST(UserLogout, a.user.Logout)
	// router.POST(UserCreate, a.user.Create)
}
