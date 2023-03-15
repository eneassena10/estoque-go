package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/* Controllers de Products */
type Controllers struct{}

func NewControllers() *Controllers {
	return &Controllers{}
}

func (c *Controllers) GetProductsAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{http.StatusOK, products})
}

func (c *Controllers) GetProductsByID(ctx *gin.Context) {
	param := ctx.Param("id")

	paramID, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	product := getProductByProductID(paramID)

	if product != nil {
		ctx.JSON(http.StatusOK, Response{http.StatusOK, product})
		return
	}

	ctx.JSON(http.StatusNotFound, Response{http.StatusNotFound, mapMessageHttp[http.StatusNotFound]})
}

func (c *Controllers) CreateProducts(ctx *gin.Context) {
	var product *Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	product.ID = products[len(products)-1].ID + 1
	products = append(products, product)

	ctx.JSON(http.StatusOK, Response{http.StatusOK, product})
}

func (c *Controllers) UpdateProducts(ctx *gin.Context) {
	paramIDProductStr := ctx.Param("id")
	paramQuantidadeStr := ctx.Param("quantidade")

	paramQuantidade, err := strconv.Atoi(paramQuantidadeStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	paramIDProduct, err := strconv.Atoi(paramIDProductStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	if paramQuantidade < 0 {
		err := errors.New("quantidade has value negative")
		ctx.JSON(http.StatusBadRequest, Response{http.StatusBadRequest, err.Error()})
		return
	}

	if p := getProductByProductID(paramIDProduct); p == nil {
		ctx.JSON(http.StatusNotFound, Response{http.StatusNotFound, mapMessageHttp[http.StatusNotFound]})
		return
	}

	var product *Product
	for _, p := range products {
		if p.ID == paramIDProduct {
			p.Quantidade += paramQuantidade
			product = p
			break
		}
	}
	ctx.JSON(http.StatusOK, Response{http.StatusOK, product})
}

func (c *Controllers) DeleteProducts(ctx *gin.Context) {
	param := ctx.Param("id")

	paramID, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}
	if p := getProductByProductID(paramID); p == nil {
		ctx.JSON(http.StatusOK, Response{http.StatusOK, mapMessageHttp[http.StatusOK]})
		return
	}

	for i, p := range products {
		if int(p.ID) == paramID {
			if p.Quantidade == 1 {
				products = append(products[:i], products[i+1:]...)
				break
			}
			p.Quantidade--
			break
		}
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func getProductByProductID(productID int) *Product {
	for _, p := range products {
		if p.ID == productID {
			return p
		}
	}
	return nil
}
