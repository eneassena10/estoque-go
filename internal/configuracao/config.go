package configuracao

import (
	"github.com/eneassena10/estoque-go/internal/controllers"
	"github.com/gin-gonic/gin"
)

const (
	// Products - rotas de listagem e criação de um novo produto
	Products = "/products"

	// ProductsID - usado em rotas que usa o id do produto para uma alteração e leitura
	ProductsList = "/products/list"
)

func InitApp(router *gin.Engine) {
	c := controllers.NewControllers()
	router.GET(Products, c.GetProductsByID)
	router.GET(ProductsList, c.GetProductsAll)
	router.DELETE(Products, c.DeleteProducts)
	router.POST(Products, c.CreateProducts)
	router.PATCH(Products, c.UpdateProductsQuantidade)
}
