package controllers

import "github.com/gin-gonic/gin"

type Product struct {
	ID         int     `json:"id,omitempty"`
	Name       string  `json:"name" badin`
	Price      float64 `json:"price"`
	Quantidade int     `json:"quantidade"`
}

type Response struct {
	Code int
	Data interface{}
}

type IControllers interface {
	GetProductsAll(ctx *gin.Context)
	GetProductsByID(ctx *gin.Context)
	CreateProducts(ctx *gin.Context)
	UpdateProductsQuantidade(ctx *gin.Context)
	DeleteProducts(ctx *gin.Context)
}

var ProductPathName = "./products.json"
