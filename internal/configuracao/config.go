package configuracao

import (
	"github.com/eneassena10/estoque-go/internal/controllers"
	"github.com/gin-gonic/gin"
)

const (
	/* Products - rotas de listagem e criação de um novo produto */
	Products = "/products"

	/* ProductsID - usado em rotas que usa o id do produto para uma alteração  */
	ProductsID = "/products/:id"

	/* ProductsIDQuatidade - altera o campo de quantidade de um produto */
	ProductsIDQuatidade = "/products/:id/:quantidade"
)

func InitApp(router *gin.Engine) {
	c := controllers.NewControllers()
	router.GET(Products, c.GetProductsAll)
	router.GET(ProductsID, c.GetProductsByID)
	router.POST(Products, c.CreateProducts)
	router.PATCH(ProductsIDQuatidade, c.UpdateProducts)
	router.DELETE(ProductsID, c.DeleteProducts)
}
