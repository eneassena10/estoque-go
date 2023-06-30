package sqlite3

import (
	"database/sql"
	"fmt"

	"github.com/eneassena10/estoque-go/internal/domain/product/entities"
	dbsqlite3 "github.com/eneassena10/estoque-go/pkg/conexao/db_sqlite3"
)

type SQLite3 struct {
	conexao *sql.DB
}

func NewSQLite3(con *sql.DB) dbsqlite3.IDataBaseOperation {
	return &SQLite3{
		conexao: con,
	}
}

func (s SQLite3) CreateEntity(entity string, data interface{}) error {
	return nil
}

func (s SQLite3) GetEntityAll(entity string) interface{} {
	query := fmt.Sprintf("SELECT id_product, name, price, quantidade FROM %s", entity)
	result, err := s.conexao.Query(query)
	if err != nil {
		return nil
	}

	products := []entities.ProductRequest{}
	for result.Next() {
		var product entities.ProductRequest
		if err := result.Scan(
			&product.ID, &product.Name, &product.Price, &product.Quantidade,
		); err != nil {
			return &[]entities.ProductRequest{}
		}

		products = append(products, product)
	}

	return products
}

func (s SQLite3) GetEntityByID(entity string, data interface{}) interface{} {
	product := data.(entities.ProductRequest)
	query := fmt.Sprintf("SELECT id_product, name, price, quantidade FROM %s WHERE id_product=?", entity)
	result, err := s.conexao.Query(query, product.ID)
	if err != nil {
		return nil
	}

	if result.Next() {
		var product entities.ProductRequest
		if err := result.Scan(
			&product.ID, &product.Name, &product.Price, &product.Quantidade,
		); err != nil {
			return nil
		}
		return product
	}
	return nil
}

func (s SQLite3) UpdateEntity(entity string, oldData interface{}, data interface{}) error {
	return nil
}

func (s SQLite3) DeleteEntity(entity string, data interface{}) error {
	return nil
}
