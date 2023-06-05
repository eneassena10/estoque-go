package controllers

import (
	"database/sql"
	"net/http"

	"github.com/eneassena10/estoque-go/internal/domain/product/entities"
	"github.com/eneassena10/estoque-go/internal/domain/product/repository"
	"github.com/eneassena10/estoque-go/internal/domain/product/service"

	"github.com/eneassena10/estoque-go/pkg/store"
	"github.com/gin-gonic/gin"
)

type IControllers interface {
	GetProductsAll(ctx *gin.Context)
	GetProductsByID(ctx *gin.Context)
	CreateProducts(ctx *gin.Context)
	UpdateProductsCount(ctx *gin.Context)
	DeleteProducts(ctx *gin.Context)
}

/*
Controllers de Products
*/
type Controllers struct {
	FileStore store.IStore
	database  *sql.DB
	Service   entities.IPoductService
}

func NewControllers(fileStore store.IStore, database *sql.DB) IControllers {
	r := repository.NewProductRepository(database)
	s := service.NewProductService(r)
	return &Controllers{
		FileStore: fileStore,
		Service:   s,
	}
}

func (c *Controllers) GetProductsAll(ctx *gin.Context) {
	// var productFileJson *entities.ProductRequest
	// if err := c.FileStore.Read(&productFileJson); err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
	// 	return
	// }
	products := c.Service.GetProductsAll(ctx)

	ctx.JSON(http.StatusOK, Response{http.StatusOK, products})
}

func (c *Controllers) GetProductsByID(ctx *gin.Context) {
	var p *entities.ProductRequest
	if err := ctx.BindJSON(&p); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	product := c.getProductByProductID(p)
	if product == nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, nil})
		return
	}

	ctx.JSON(http.StatusOK, Response{http.StatusOK, product})
}

func (c *Controllers) CreateProducts(ctx *gin.Context) {
	var productRequest *entities.ProductRequest
	err := ctx.ShouldBindJSON(&productRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	var productList []*entities.ProductRequest
	err = c.FileStore.Read(&productList)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	if len(productList) > 0 {
		productRequest.ID = productList[len(productList)-1].ID + 1
	} else {
		productRequest.ID = 1
	}
	productList = append(productList, productRequest)
	err = c.FileStore.Write(&productList)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Response{http.StatusOK, productRequest})
}

func (c *Controllers) UpdateProductsCount(ctx *gin.Context) {
	var product *entities.ProductRequest
	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	var productList []*entities.ProductRequest
	err := c.FileStore.Read(&productList)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	for _, p := range productList {
		if p.ID == product.ID && (p.Quantidade+product.Quantidade) >= 0 {
			p.Quantidade += product.Quantidade
			product = p
			break
		}
	}

	if err := c.FileStore.Write(&productList); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Response{http.StatusOK, product})
}

func (c *Controllers) DeleteProducts(ctx *gin.Context) {
	var product *entities.ProductRequest

	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}
	pList, err := c.loadListProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	for i, p := range pList {
		if p.ID == product.ID {
			pList = append(pList[:i], pList[i+1:]...)
		}
	}

	err = c.FileStore.Write(&pList)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{http.StatusInternalServerError, err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (c *Controllers) getProductByProductID(product *entities.ProductRequest) *entities.ProductRequest {
	products, _ := c.loadListProducts()
	for _, p := range products {
		if p.ID == product.ID {
			return p
		}
	}
	return nil
}

func (c *Controllers) loadListProducts() ([]*entities.ProductRequest, error) {
	var p []*entities.ProductRequest
	if err := c.FileStore.Read(&p); err != nil {
		return p, err
	}
	return p, nil
}
