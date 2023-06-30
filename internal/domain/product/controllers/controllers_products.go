package controllers

import (
	"net/http"

	"github.com/eneassena10/estoque-go/internal/domain/product/entities"
	"github.com/eneassena10/estoque-go/internal/domain/product/service"
	"github.com/eneassena10/estoque-go/pkg/regras"
	"github.com/gin-gonic/gin"
)

type ProductControllers struct {
	Service *service.ProductService
}

func NewControllers(database *service.ProductService) *ProductControllers {
	return &ProductControllers{
		Service: database,
	}
}

func (c *ProductControllers) GetProductsAll(ctx *gin.Context) {
	products := c.Service.GetProductsAll()

	ctx.JSON(http.StatusOK, Response{http.StatusOK, products})
}

func (c *ProductControllers) GetProductsByID(ctx *gin.Context) {
	var requestBody ProductRequestBody

	if regras.ValidateErrorInRequest(ctx, &requestBody) {
		return
	}
	productSearch := entities.NewProduct().WithID(requestBody.ID)
	product := c.Service.GetProductsOne(productSearch)
	if product == nil {
		ctx.JSON(http.StatusBadRequest, Response{http.StatusBadRequest, requestBody})
		return
	}

	ctx.JSON(http.StatusOK, Response{http.StatusOK, product})
}

func (c *ProductControllers) CreateProducts(ctx *gin.Context) {
	var requestBody entities.ProductRequest
	if regras.ValidateErrorInRequest(ctx, &requestBody) {
		return
	}
	productCreated := entities.NewProduct().
		WithName(requestBody.Name).
		WithPrice(requestBody.Price).
		WithQuantidade(requestBody.Quantidade)

	if err := c.Service.CreateProducts(productCreated); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Response{http.StatusOK, requestBody})
}

func (c *ProductControllers) UpdateProductsCount(ctx *gin.Context) {
	var requestBody *ProductRequestBody
	if regras.ValidateErrorInRequest(ctx, &requestBody) {
		return
	}

	productUpdateCount := entities.NewProduct().
		WithID(requestBody.ID).
		WithQuantidade(requestBody.Quantidade)

	if err := c.Service.UpdateProductsCount(productUpdateCount); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}
	productResponse := c.Service.GetProductsOne(productUpdateCount)
	ctx.JSON(http.StatusOK, Response{http.StatusOK, productResponse})
}

func (c *ProductControllers) DeleteProducts(ctx *gin.Context) {
	var requestBody *ProductRequestBody
	if regras.ValidateErrorInRequest(ctx, &requestBody) {
		return
	}
	productDeleted := entities.NewProduct().WithID(requestBody.ID)
	if err := c.Service.DeleteProducts(productDeleted); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
