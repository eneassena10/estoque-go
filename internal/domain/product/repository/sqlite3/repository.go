package sqlite3_repository

import (
	"database/sql"

	"github.com/eneassena10/estoque-go/internal/domain/product/entities"
)

//go:generate mockgen -source=./repository.go -destination=./../../../test/mockgen/product_repository_mock.go -package=mockgen
type IProductRepository interface {
	GetProductsAll() *[]entities.ProductRequest
	GetProductsOne(product *entities.ProductRequest) *entities.ProductRequest
	CreateProducts(product *entities.ProductRequest) error
	UpdateProductsCount(oldProduct *entities.ProductRequest, product *entities.ProductRequest) error
	DeleteProducts(product *entities.ProductRequest) error
}

type ProductRepository struct {
	dataBase *sql.DB
}

func NewProductRepository(database *sql.DB) IProductRepository {
	return &ProductRepository{dataBase: database}
}

func (r *ProductRepository) GetProductsAll() *[]entities.ProductRequest {
	result, err := r.dataBase.Query(QUERY_SELECT_ALL_PRODUCT)
	if err != nil {
		return &[]entities.ProductRequest{}
	}

	defer result.Close()

	products := &[]entities.ProductRequest{}
	for result.Next() {

		var product entities.ProductRequest
		if err := result.Scan(
			&product.ID, &product.Name, &product.Price, &product.Quantidade,
		); err != nil {
			return &[]entities.ProductRequest{}
		}

		*products = append(*products, product)
	}
	return products
}

func (r *ProductRepository) GetProductsOne(product *entities.ProductRequest) *entities.ProductRequest {
	stmt, err := r.dataBase.Query(QUERY_SELECT_BY_ID_PRODUCT, &product.ID)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	if stmt.Next() {
		var p entities.ProductRequest
		if err := stmt.Scan(&p.ID, &p.Name, &p.Price, &p.Quantidade); err != nil {
			return nil
		}
		return &p
	}
	return nil
}

func (r *ProductRepository) CreateProducts(product *entities.ProductRequest) error {
	stmt, err := r.dataBase.Prepare(QUERY_CREATE_PRODUCT)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(&product.Name, &product.Price, &product.Quantidade)
	if err != nil {
		return err
	}
	if rowsAffected, err := result.RowsAffected(); err != nil && rowsAffected == 0 {
		return err
	}
	return nil
}

func (r *ProductRepository) UpdateProductsCount(oldProduct, product *entities.ProductRequest) error {
	// dbProduct := r.GetProductsOne(product)
	if oldProduct != nil && oldProduct.Quantidade+product.Quantidade >= 0 {
		oldProduct.Quantidade = oldProduct.Quantidade + product.Quantidade
		result, err := r.dataBase.Exec(QUERY_UPDATE_COUNT_PRODUCT, &oldProduct.Quantidade, &oldProduct.ID)
		if err != nil {
			return err
		}
		if rowsAffected, err := result.RowsAffected(); err != nil && rowsAffected == 0 {
			return err
		}
	}
	return nil
}

func (r *ProductRepository) DeleteProducts(product *entities.ProductRequest) error {
	result, err := r.dataBase.Exec(QUERY_DELETE_PRODUCT, &product.ID)
	if err != nil {
		return err
	}

	if rowsAffected, err := result.RowsAffected(); err != nil && rowsAffected == 0 {
		return err
	}
	return nil
}
