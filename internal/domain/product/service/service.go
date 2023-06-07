package service

import (
	"github.com/gin-gonic/gin"

	"github.com/eneassena10/estoque-go/internal/domain/product/entities"
)

type ProductService struct {
	Repository entities.IProductRepository
}

func NewProductService(repository entities.IProductRepository) entities.IPoductService {
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
	return s.Repository.CreateProducts(ctx, product)
}

func (s *ProductService) UpdateProductsCount(ctx *gin.Context, product *entities.ProductRequest) error {
	productSearch := s.Repository.GetProductsOne(ctx, product)
	return s.Repository.UpdateProductsCount(ctx, productSearch, product)
}

func (s *ProductService) DeleteProducts(ctx *gin.Context, product *entities.ProductRequest) error {
	return nil
}
