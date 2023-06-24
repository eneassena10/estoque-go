package sqlite3_repository

import (
	"database/sql"

	"github.com/eneassena10/estoque-go/internal/domain/product/entities"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=./repository.go -destination=./../../../test/mockgen/product_repository_mock.go -package=mockgen
type IProductRepository interface {
	GetProductsAll(ctx *gin.Context) *[]entities.ProductRequest
	GetProductsOne(ctx *gin.Context, product *entities.ProductRequest) *entities.ProductRequest
	CreateProducts(ctx *gin.Context, product *entities.ProductRequest) error
	UpdateProductsCount(ctx *gin.Context, oldProduct *entities.ProductRequest, product *entities.ProductRequest) error
	DeleteProducts(ctx *gin.Context, product *entities.ProductRequest) error
}

type ProductRepository struct {
	dataBase *sql.DB
}

func NewProductRepository(database *sql.DB) IProductRepository {
	return &ProductRepository{dataBase: database}
}

func (r *ProductRepository) GetProductsAll(ctx *gin.Context) *[]entities.ProductRequest {
	result, err := r.dataBase.QueryContext(ctx, QUERY_SELECT_ALL_PRODUCT)
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

func (r *ProductRepository) GetProductsOne(ctx *gin.Context, product *entities.ProductRequest) *entities.ProductRequest {
	stmt, err := r.dataBase.QueryContext(ctx, QUERY_SELECT_BY_ID_PRODUCT, &product.ID)
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

func (r *ProductRepository) CreateProducts(ctx *gin.Context, product *entities.ProductRequest) error {
	stmt, err := r.dataBase.PrepareContext(ctx, QUERY_CREATE_PRODUCT)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(ctx, &product.Name, &product.Price, &product.Quantidade)
	if err != nil {
		return err
	}
	if rowsAffected, err := result.RowsAffected(); err != nil && rowsAffected == 0 {
		return err
	}
	return nil
}

func (r *ProductRepository) UpdateProductsCount(ctx *gin.Context, oldProduct, product *entities.ProductRequest) error {
	dbProduct := r.GetProductsOne(ctx, product)
	if dbProduct != nil && dbProduct.Quantidade+product.Quantidade >= 0 {
		dbProduct.Quantidade = dbProduct.Quantidade + product.Quantidade
		result, err := r.dataBase.ExecContext(ctx, QUERY_UPDATE_COUNT_PRODUCT, &dbProduct.Quantidade, &dbProduct.ID)
		if err != nil {
			return err
		}
		if rowsAffected, err := result.RowsAffected(); err != nil && rowsAffected == 0 {
			return err
		}
	}
	return nil
}

func (r *ProductRepository) DeleteProducts(ctx *gin.Context, product *entities.ProductRequest) error {
	result, err := r.dataBase.ExecContext(ctx, QUERY_DELETE_PRODUCT, &product.ID)
	if err != nil {
		return err
	}
	if rowsAffected, err := result.RowsAffected(); err != nil && rowsAffected == 0 {
		return err
	}
	return nil
}
