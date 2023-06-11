package sqlite3_repository

import (
	"database/sql"

	"github.com/eneassena10/estoque-go/internal/domain/product/entities"
	"github.com/gin-gonic/gin"
)

type ProductRepository struct {
	dataBase *sql.DB
}

const (
	QUERY_SELECT_ALL_PRODUCT   = "SELECT id_product, name, price, quantidade FROM products;"
	QUERY_SELECT_BY_ID_PRODUCT = "SELECT id_product, name, price, quantidade FROM products WHERE id_product=?;"
	QUERY_UPDATE_COUNT_PRODUCT = "UPDATE products SET quantidade=? WHERE id_product=?;"
	QUERY_DELETE_PRODUCT       = "DELETE FROM products WHERE id_product=?;"
	QUERY_CREATE_PRODUCT       = "INSERT INTO products (name, price, quantidade) VALUES(?, ?, ?);"
)

func NewProductRepository(database *sql.DB) entities.IProductRepository {
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
