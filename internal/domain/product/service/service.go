package service

import (
	"github.com/eneassena10/estoque-go/internal/domain/product/entities"
	sqlite3_repository "github.com/eneassena10/estoque-go/internal/domain/product/repository"
)

type ProductService struct {
	Repository *sqlite3_repository.ProductRepository
}

func NewProductService(repository *sqlite3_repository.ProductRepository) *ProductService {
	return &ProductService{Repository: repository}
}

func (s *ProductService) GetProductsAll() *[]entities.ProductRequest {
	products := s.Repository.GetProductsAll()
	return products
}

func (s *ProductService) GetProductsOne(product *entities.ProductRequest) *entities.ProductRequest {
	p := s.Repository.GetProductsOne(product)
	return p
}

func (s *ProductService) CreateProducts(product *entities.ProductRequest) error {
	erroCreate := s.Repository.CreateProducts(product)
	if erroCreate != nil {
		return erroCreate
	}
	return nil
}

func (s *ProductService) UpdateProductsCount(product *entities.ProductRequest) error {
	// productSearch := s.GetProductsOne(product)
	// if productSearch == nil {
	// 	return errors.New("not found product")
	// }
	// return s.Repository.UpdateProductsCount(productSearch, product)
	return nil
}

func (s *ProductService) DeleteProducts(product *entities.ProductRequest) error {
	erroDelete := s.Repository.DeleteProducts(product)
	if erroDelete != nil {
		return erroDelete
	}
	return nil
}
