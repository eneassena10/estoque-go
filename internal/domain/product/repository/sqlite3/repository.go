package sqlite3

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
	return &[]entities.ProductRequest{}
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
