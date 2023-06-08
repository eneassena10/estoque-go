package sqlite3_repository

import (
	"database/sql"

	"github.com/eneassena10/estoque-go/internal/domain/product/entities"
	"github.com/gin-gonic/gin"
)

type ProductRepository struct {
	dataBase *sql.DB
}

func NewProductRepository(database *sql.DB) entities.IProductRepository {
	return &ProductRepository{dataBase: database}
}

func (r *ProductRepository) GetProductsAll(ctx *gin.Context) *[]entities.ProductRequest {
	result, err := r.dataBase.QueryContext(ctx, "select id_product, name,price,quantidade from products;")
	if err != nil {
		return &[]entities.ProductRequest{}
	}

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
	return &entities.ProductRequest{}
}

func (r *ProductRepository) CreateProducts(ctx *gin.Context, product *entities.ProductRequest) error {
	return nil
}

func (r *ProductRepository) UpdateProductsCount(ctx *gin.Context, oldProduct, product *entities.ProductRequest) error {
	return nil
}

func (r *ProductRepository) DeleteProducts(ctx *gin.Context, product *entities.ProductRequest) error {
	return nil
}
