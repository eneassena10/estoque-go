package controllers

import (
	"database/sql"
	"net/http"

	"github.com/eneassena10/estoque-go/internal/domain/product/entities"
	sqlite3_repository "github.com/eneassena10/estoque-go/internal/domain/product/repository/sqlite3"
	"github.com/eneassena10/estoque-go/internal/domain/product/service"

	"github.com/gin-gonic/gin"
)

type IProductControllers interface {
	GetProductsAll(ctx *gin.Context)
	GetProductsByID(ctx *gin.Context)
	CreateProducts(ctx *gin.Context)
	UpdateProductsCount(ctx *gin.Context)
	DeleteProducts(ctx *gin.Context)
}

/*
Controllers de Products
*/
type ProductControllers struct {
	Service entities.IPoductService
}

func NewControllers(database *sql.DB) IProductControllers {
	r := sqlite3_repository.NewProductRepository(database)
	s := service.NewProductService(r)
	return &ProductControllers{
		Service: s,
	}
}

func (c *ProductControllers) GetProductsAll(ctx *gin.Context) {
	products := c.Service.GetProductsAll(ctx)

	ctx.JSON(http.StatusOK, Response{http.StatusOK, products})
}

func (c *ProductControllers) GetProductsByID(ctx *gin.Context) {
	var requestBody ProductRequestBody
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}
	productSearch := &entities.ProductRequest{ID: requestBody.ID}
	product := c.Service.GetProductsOne(ctx, productSearch)
	if product == nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, nil})
		return
	}

	ctx.JSON(http.StatusOK, Response{http.StatusOK, product})
}

func (c *ProductControllers) CreateProducts(ctx *gin.Context) {
	var requestBody entities.ProductRequest
	err := ctx.BindJSON(&requestBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	if err = c.Service.CreateProducts(ctx, &requestBody); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Response{http.StatusOK, requestBody})
}

func (c *ProductControllers) UpdateProductsCount(ctx *gin.Context) {
	var requestBody *entities.ProductRequest
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	if err := c.Service.UpdateProductsCount(ctx, requestBody); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Response{http.StatusOK, requestBody})
}

func (c *ProductControllers) DeleteProducts(ctx *gin.Context) {
	var product *entities.ProductRequest

	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}
	if err := c.Service.DeleteProducts(ctx, product); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
