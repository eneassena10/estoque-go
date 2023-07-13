package sqlite3

import (
	"database/sql"
	"fmt"
	"strings"

	dbsqlite3 "github.com/eneassena10/estoque-go/pkg/conexao/db_sqlite3"
)

type SQLite3 struct {
	conexao   *sql.DB
	resultSet interface{}
}

const (
	SELECT = "SELECT"
	FROM   = "FROM"
	WHERE  = "WHERE"
	DELETE = "DELETE"
	UPDATE = "UPDADE"
	TABLE  = "TABLE"
	SET    = "SET"
)

func NewSQLite3(con *sql.DB) dbsqlite3.IDataBaseOperation {
	return &SQLite3{
		conexao: con,
	}
}

func (s SQLite3) CreateEntity(entity string, data interface{}) error {
	return nil
}

func (s SQLite3) GetEntityAll(entity string, fields []string) interface{} {
	//  entity, fields, estrutura
	dataFields := strings.Join(fields, ",")
	dataQuery := fmt.Sprintf("%s %s %s %s", SELECT, dataFields, FROM, entity)

	result, err := s.conexao.Query(dataQuery)
	if err != nil {
		return nil
	}
	s.resultSet = result
	return s.resultSet
}

func (s SQLite3) GetEntityByID(entity string, fields []string, data interface{}) interface{} {
	dataFields := strings.Join(fields, ",")
	entityID := data.(int)
	query := fmt.Sprintf("%s %s %s %s %s %s=?", SELECT, dataFields, FROM, entity, WHERE, dataFields[0])
	result, err := s.conexao.Query(query, entityID)
	if err != nil {
		return nil
	}
	s.resultSet = result

	// if result.Next() {
	// 	var product entities.ProductRequest
	// 	if err := result.Scan(
	// 		&product.ID, &product.Name, &product.Price, &product.Quantidade,
	// 	); err != nil {
	// 		return nil
	// 	}
	// 	return product
	// }
	return s.resultSet
}

func (s SQLite3) UpdateEntity(entity string, oldData interface{}, data interface{}) error {
	return nil
}

func (s SQLite3) DeleteEntity(entity string, data interface{}) error {
	return nil
}
