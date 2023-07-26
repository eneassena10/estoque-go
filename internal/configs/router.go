package configs

import (
	pControllers "github.com/eneassena10/estoque-go/internal/domain/product/controllers"
	uControllers "github.com/eneassena10/estoque-go/internal/domain/user/controllers"
	"github.com/gin-gonic/gin"
)

var (
	// Products - rotas de listagem e criação de um novo produto
	Products = "/products"

	// ProductsID - usado em rotas que usa o id do produto para uma alteração e leitura
	ProductsList = "/products/list"
	UserLogin    = "/user/login"
	UserLogout   = "/user/logout"
	UserCreate   = "/user/create"
)

type App struct {
	Products *pControllers.ProductControllers
	Users    *uControllers.UserController
}

func NewApp(mapServices map[HandleNameType]interface{}) *App {
	app := &App{}
	app.handlers(mapServices)
	return app
}

func (a *App) handlers(mapServices map[HandleNameType]interface{}) {
	a.Products = mapServices[HandleProduct].(*pControllers.ProductControllers)
	a.Users = nil
}

func (a *App) Routers(router *gin.Engine) {
	router.GET(ProductsList, a.Products.GetProductsAll)
	router.GET(Products, a.Products.GetProductsByID)
	router.DELETE(Products, a.Products.DeleteProducts)
	router.POST(Products, a.Products.CreateProducts)
	router.PATCH(Products, a.Products.UpdateProductsCount)

	// router.POST(UserLogin, a.users.Logar)
	// router.POST(UserLogout, a.users.Logout)
	// router.POST(UserCreate, a.users.Create)
}
