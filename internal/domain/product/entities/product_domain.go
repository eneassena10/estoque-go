package entities

import "github.com/gin-gonic/gin"

type Product struct {
	ID         int
	Name       string
	Price      float64
	Quantidade int
}

type ProductRequest struct {
	ID         int     `json:"id,omitempty"`
	Name       string  `json:"name" binding:"requered"`
	Price      float64 `json:"price" binding:"requered"`
	Quantidade int     `json:"quantidade" binding:"requered"`
}

var ProductPathName = "./products.json"

type IProductRepository interface {
	GetProductsAll(ctx *gin.Context) *[]ProductRequest
	GetProductsOne(ctx *gin.Context, product *ProductRequest) *ProductRequest
	CreateProducts(ctx *gin.Context, product *ProductRequest) error
	UpdateProductsCount(ctx *gin.Context, product *ProductRequest) error
	DeleteProducts(ctx *gin.Context, product *ProductRequest) error
}

type IPoductService interface {
	GetProductsAll(ctx *gin.Context) *[]ProductRequest
	GetProductsOne(ctx *gin.Context, product *ProductRequest) *ProductRequest
	CreateProducts(ctx *gin.Context, product *ProductRequest) error
	UpdateProductsCount(ctx *gin.Context, oldProduct, product *ProductRequest) error
	DeleteProducts(ctx *gin.Context, product *ProductRequest) error
}
