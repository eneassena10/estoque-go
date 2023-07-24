package sqlite3_repository

import (
	"database/sql"

	"github.com/eneassena10/estoque-go/internal/domain/product/entities"
	dbsqlite3 "github.com/eneassena10/estoque-go/pkg/conexao/db_sqlite3"
)

type ProductRepository struct {
	operationSql dbsqlite3.IDataBaseOperation
}

const EntityName string = "products"

func NewProductRepository(operationSql dbsqlite3.IDataBaseOperation) *ProductRepository {
	return &ProductRepository{operationSql: operationSql}
}

func (r *ProductRepository) GetProductsAll() *[]entities.ProductRequest {
	fields := []string{"id_product", "name", "price", "count"}
	result := r.operationSql.GetEntityAll(EntityName, fields)
	if result != nil {
		return r.getProductsAll(result)
	}
	return &[]entities.ProductRequest{}
}

func (r ProductRepository) getProductsAll(result interface{}) *[]entities.ProductRequest {
	resultQuery := result.(*sql.Rows)
	resultSet := []entities.ProductRequest{}
	for resultQuery.Next() {
		var product entities.ProductRequest
		if err := resultQuery.Scan(&product.ID, &product.Name, &product.Price, &product.Count); err != nil {
			return &[]entities.ProductRequest{}
		}
		resultSet = append(resultSet, product)
	}
	return &resultSet
}

func (r *ProductRepository) GetProductsOne(product *entities.ProductRequest) *entities.ProductRequest {
	fields := []string{"id_product", "name", "price", "count"}
	if result := r.operationSql.GetEntityByID(EntityName, fields, product); result != nil {
		r := result.(*sql.Row)
		return getProductsOne(r)
	}
	return nil
}

func getProductsOne(result *sql.Row) *entities.ProductRequest {
	var product entities.ProductRequest
	if err := result.Scan(&product.ID, &product.Name, &product.Price, &product.Count); err != nil {
		return nil
	}
	return &product
}

func (r *ProductRepository) CreateProducts(product *entities.ProductRequest) error {
	fields := []string{"name", "price", "count"}
	err := r.operationSql.CreateEntity(EntityName, fields, product)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) UpdateProductsCount(oldProduct, product *entities.ProductRequest) error {
	// if oldProduct != nil && oldProduct.Quantidade+product.Quantidade >= 0 {
	// 	oldProduct.Quantidade = oldProduct.Quantidade + product.Quantidade
	// 	result, err := r.dataBase.Exec(QUERY_UPDATE_COUNT_PRODUCT, &oldProduct.Quantidade, &oldProduct.ID)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if rowsAffected, err := result.RowsAffected(); err != nil && rowsAffected == 0 {
	// 		return err
	// 	}
	// }
	return nil
}

func (r *ProductRepository) DeleteProducts(product *entities.ProductRequest) error {
	fields := []string{"id_product"}
	if err := r.operationSql.DeleteEntity(EntityName, fields, product); err != nil {
		return err
	}
	return nil
}
