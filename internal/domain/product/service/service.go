package service

import (
	"errors"

	"github.com/eneassena10/estoque-go/internal/domain/product/entities"
	sqlite3_repository "github.com/eneassena10/estoque-go/internal/domain/product/repository/sqlite3"
)

//go:generate mockgen -source=./service.go -destination=./../../test/mockgen/product_service_mock.go -package=mockgen
type IPoductService interface {
	GetProductsAll() *[]entities.ProductRequest
	GetProductsOne(product *entities.ProductRequest) *entities.ProductRequest
	CreateProducts(product *entities.ProductRequest) error
	UpdateProductsCount(oldProduct *entities.ProductRequest) error
	DeleteProducts(product *entities.ProductRequest) error
}
type ProductService struct {
	Repository sqlite3_repository.IProductRepository
}

func NewProductService(repository sqlite3_repository.IProductRepository) IPoductService {
	return &ProductService{Repository: repository}
}

func (s *ProductService) GetProductsAll() *[]entities.ProductRequest {
	p := s.Repository.GetProductsAll()
	return p
}

func (s *ProductService) GetProductsOne(product *entities.ProductRequest) *entities.ProductRequest {
	p := s.Repository.GetProductsOne(product)
	return p
}

func (s *ProductService) CreateProducts(product *entities.ProductRequest) error {
	erroCreate := s.Repository.CreateProducts(product)
	if erroCreate == nil {
		return erroCreate
	}
	return errors.New("não possível criar um novo produto! 'Tente Novamente'")
}

func (s *ProductService) UpdateProductsCount(product *entities.ProductRequest) error {
	productSearch := s.GetProductsOne(product)
	if productSearch == nil {
		return errors.New("not found product")
	}
	return s.Repository.UpdateProductsCount(productSearch, product)
}

func (s *ProductService) DeleteProducts(product *entities.ProductRequest) error {
	err := s.Repository.DeleteProducts(product)
	return err
}
