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
	return &[]entities.ProductRequest{}
}

func (s *ProductService) GetProductsOne(ctx *gin.Context, product *entities.ProductRequest) *entities.ProductRequest {
	return &entities.ProductRequest{}
}

func (s *ProductService) CreateProducts(ctx *gin.Context, product *entities.ProductRequest) error {
	return nil
}

func (s *ProductService) UpdateProductsCount(ctx *gin.Context, oldProduct, product *entities.ProductRequest) error {
	return nil
}

func (s *ProductService) DeleteProducts(ctx *gin.Context, product *entities.ProductRequest) error {
	return nil
}
