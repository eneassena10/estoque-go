package sqlite3_repository

import (
	"database/sql"
	"log"

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
	query := "SELECT id_product, name,price,quantidade FROM products where id_product = ?;"
	stmt, err := r.dataBase.QueryContext(ctx, query, &product.ID)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	if stmt.Next() {
		var p entities.ProductRequest
		_ = stmt.Scan(&p.ID, &p.Name, &p.Price, &p.Quantidade)
		return &p
	}

	return nil
}

func (r *ProductRepository) CreateProducts(ctx *gin.Context, product *entities.ProductRequest) error {
	query := "INSERT INTO products (name, price, quantidade) VALUES(?,?,?)"
	stmt, err := r.dataBase.PrepareContext(ctx, query)
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

func (r *ProductRepository) UpdateProductsCount(ctx *gin.Context, oldProduct, product *entities.ProductRequest) error {
	query := "UPDATE products SET name=?, price=?, quantidade=? WHERE id_product=?;"
	result, err := r.dataBase.Exec(query, &product.Name, &product.Price, &product.Quantidade, &product.ID)
	if err != nil {
		return err
	}
	if rowsAffected, err := result.RowsAffected(); err != nil && rowsAffected == 0 {
		log.Printf("rowsAffected: [%v]", rowsAffected)
		return err
	}
	return nil
}

func (r *ProductRepository) DeleteProducts(ctx *gin.Context, product *entities.ProductRequest) error {
	query := "DELETE FROM products where id_product=3;"
	result, err := r.dataBase.Exec(query, &product.ID)
	if err != nil {
		return err
	}
	if rowsAffected, err := result.RowsAffected(); err != nil && rowsAffected == 0 {
		log.Printf("rowsAffected: [%v]", rowsAffected)
		return err
	}
	return nil
}
