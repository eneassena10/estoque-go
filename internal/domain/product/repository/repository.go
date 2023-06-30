package sqlite3_repository

import (
	"github.com/eneassena10/estoque-go/internal/domain/product/entities"
	dbsqlite3 "github.com/eneassena10/estoque-go/pkg/conexao/db_sqlite3"
)

type ProductRepository struct {
	operationSql dbsqlite3.IDataBaseOperation
}

func NewProductRepository(operationSql interface{}) *ProductRepository {
	opSql := operationSql.(dbsqlite3.IDataBaseOperation)
	return &ProductRepository{operationSql: opSql}
}

func (r *ProductRepository) GetProductsAll() *[]entities.ProductRequest {
	if products := r.operationSql.GetEntityAll("products"); products != nil {
		productsAll := products.([]entities.ProductRequest)
		return &productsAll
	}
	return &[]entities.ProductRequest{}
}

func (r *ProductRepository) GetProductsOne(product *entities.ProductRequest) *entities.ProductRequest {
	if resultProduct := r.operationSql.GetEntityByID("products", *product); resultProduct != nil {
		p := resultProduct.(entities.ProductRequest)
		return &p
	}
	return nil
}

func (r *ProductRepository) CreateProducts(product *entities.ProductRequest) error {
	// stmt, err := r.dataBase.Prepare(QUERY_CREATE_PRODUCT)
	// if err != nil {
	// 	return err
	// }
	// defer stmt.Close()
	// result, err := stmt.Exec(&product.Name, &product.Price, &product.Quantidade)
	// if err != nil {
	// 	return err
	// }
	// if rowsAffected, err := result.RowsAffected(); err != nil && rowsAffected == 0 {
	// 	return err
	// }
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
	// result, err := r.dataBase.Exec(QUERY_DELETE_PRODUCT, &product.ID)
	// if err != nil {
	// 	return err
	// }

	// if rowsAffected, err := result.RowsAffected(); err != nil && rowsAffected == 0 {
	// 	return err
	// }
	return nil
}
