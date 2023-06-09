package service

import (
	"errors"
	"log"

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
	if newProduct := s.Repository.CreateProducts(ctx, product); newProduct == nil {
		return newProduct
	}
	return errors.New("não possível criar um novo produto! 'Tente Novamente'")
}

func (s *ProductService) UpdateProductsCount(ctx *gin.Context, product *entities.ProductRequest) error {
	productSearch := s.GetProductsOne(ctx, product)
	log.Println(productSearch)
	if productSearch == nil {
		return errors.New("not found product")
	}
	return s.Repository.UpdateProductsCount(ctx, productSearch, product)
}

func (s *ProductService) DeleteProducts(ctx *gin.Context, product *entities.ProductRequest) error {
	err := s.Repository.DeleteProducts(ctx, product)
	return err
}
