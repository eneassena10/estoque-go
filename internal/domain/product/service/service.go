package service

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/eneassena10/estoque-go/internal/domain/product/entities"
	sqlite3_repository "github.com/eneassena10/estoque-go/internal/domain/product/repository/sqlite3"
)

//go:generate mockgen -source=./service.go -destination=./../../test/mockgen/product_service_mock.go -package=mockgen
type IPoductService interface {
	GetProductsAll(ctx *gin.Context) *[]entities.ProductRequest
	GetProductsOne(ctx *gin.Context, product *entities.ProductRequest) *entities.ProductRequest
	CreateProducts(ctx *gin.Context, product *entities.ProductRequest) error
	UpdateProductsCount(ctx *gin.Context, oldProduct *entities.ProductRequest) error
	DeleteProducts(ctx *gin.Context, product *entities.ProductRequest) error
}
type ProductService struct {
	Repository sqlite3_repository.IProductRepository
}

func NewProductService(repository sqlite3_repository.IProductRepository) IPoductService {
	return &ProductService{Repository: repository}
}

func (s *ProductService) GetProductsAll(ctx *gin.Context) *[]entities.ProductRequest {
	p := s.Repository.GetProductsAll(ctx)
	return p
}

func (s *ProductService) GetProductsOne(ctx *gin.Context, product *entities.ProductRequest) *entities.ProductRequest {
	p := s.Repository.GetProductsOne(ctx, product)
	return p
}

func (s *ProductService) CreateProducts(ctx *gin.Context, product *entities.ProductRequest) error {
	erroCreate := s.Repository.CreateProducts(ctx, product)
	if erroCreate == nil {
		return erroCreate
	}
	return errors.New("não possível criar um novo produto! 'Tente Novamente'")
}

func (s *ProductService) UpdateProductsCount(ctx *gin.Context, product *entities.ProductRequest) error {
	productSearch := s.GetProductsOne(ctx, product)
	if productSearch == nil {
		return errors.New("not found product")
	}
	return s.Repository.UpdateProductsCount(ctx, productSearch, product)
}

func (s *ProductService) DeleteProducts(ctx *gin.Context, product *entities.ProductRequest) error {
	err := s.Repository.DeleteProducts(ctx, product)
	return err
}
