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
	Name       string  `json:"name" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
	Quantidade int     `json:"quantidade" binding:"required"`
}

type IProductRepository interface {
	GetProductsAll(ctx *gin.Context) *[]ProductRequest
	GetProductsOne(ctx *gin.Context, product *ProductRequest) *ProductRequest
	CreateProducts(ctx *gin.Context, product *ProductRequest) error
	UpdateProductsCount(ctx *gin.Context, oldProduct *ProductRequest, product *ProductRequest) error
	DeleteProducts(ctx *gin.Context, product *ProductRequest) error
}

type IPoductService interface {
	GetProductsAll(ctx *gin.Context) *[]ProductRequest
	GetProductsOne(ctx *gin.Context, product *ProductRequest) *ProductRequest
	CreateProducts(ctx *gin.Context, product *ProductRequest) error
	UpdateProductsCount(ctx *gin.Context, oldProduct *ProductRequest) error
	DeleteProducts(ctx *gin.Context, product *ProductRequest) error
}

func NewProduct() *ProductRequest {
	return &ProductRequest{}
}

func (p *ProductRequest) WithID(id int) *ProductRequest {
	if id != 0 {
		p.ID = id
	}
	return p
}

func (p *ProductRequest) WithName(name string) *ProductRequest {
	if name != "" {
		p.Name = name
	}
	return p
}

func (p *ProductRequest) WithPrice(price float64) *ProductRequest {
	if price != 0 {
		p.Price = price
	}
	return p
}

func (p *ProductRequest) WithQuantidade(quantidade int) *ProductRequest {
	if quantidade != 0 {
		p.Quantidade = quantidade
	}
	return p
}
