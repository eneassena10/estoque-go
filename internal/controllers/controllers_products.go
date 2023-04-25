package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Controllers de Products
*/
type Controllers struct{}

func NewControllers() IControllers {
	return &Controllers{}
}

func (c *Controllers) GetProductsAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{http.StatusOK, products})
}

func (c *Controllers) GetProductsByID(ctx *gin.Context) {
	var idProduct Product
	if err := ctx.BindJSON(&idProduct); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	product := c.getProductByProductID(idProduct.ID)
	if product == nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, nil})
		return
	}

	ctx.JSON(http.StatusOK, Response{http.StatusOK, product})
}

func (c *Controllers) CreateProducts(ctx *gin.Context) {
	var product *Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}
	if len(products) > 0 {
		product.ID = products[len(products)-1].ID + 1
	} else {
		product.ID = 1
	}

	products = append(products, product)

	ctx.JSON(http.StatusOK, Response{http.StatusOK, product})
}

func (c *Controllers) UpdateProductsQuantidade(ctx *gin.Context) {
	var productTheRequest *Product
	if err := ctx.BindJSON(&productTheRequest); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	product, err := c.updateParcialProduct(productTheRequest)
	if product == nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Response{http.StatusOK, product})
}

func (c *Controllers) DeleteProducts(ctx *gin.Context) {
	var product *Product

	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	for i, p := range products {
		if int(p.ID) == product.ID {
			products = append(products[:i], products[i+1:]...)
			break
		}
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (c *Controllers) getProductByProductID(productID int) *Product {
	for _, p := range products {
		if p.ID == productID {
			return p
		}
	}
	return nil
}

func (c *Controllers) updateParcialProduct(p *Product) (*Product, error) {
	product := c.getProductByProductID(p.ID)
	if product == nil {
		return &Product{}, errors.New("product not found")
	}

	if (product.Quantidade + p.Quantidade) >= 0 {
		product.Quantidade += p.Quantidade
	}

	return product, nil
}
