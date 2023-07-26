package sqlite3_repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/eneassena10/estoque-go/internal/domain/product/entities"
	dbsqlite3 "github.com/eneassena10/estoque-go/pkg/conexao/db_sqlite3"
)

type SQLite3Repository struct {
	conexao   *sql.DB
	resultSet interface{}
}

const (
	INSERT = "INSERT"
	INTO   = "INTO"
	VALUES = "VALUES"
	SELECT = "SELECT"
	FROM   = "FROM"
	WHERE  = "WHERE"
	DELETE = "DELETE"
	UPDATE = "UPDATE"
	TABLE  = "TABLE"
	SET    = "SET"
)

func NewSQLite3(con *sql.DB) dbsqlite3.IDataBaseOperation {
	return &SQLite3Repository{
		conexao: con,
	}
}

func (s SQLite3Repository) CreateEntity(entity string, fields []string, data interface{}) error {
	dataFields := strings.Join(fields, ",")
	product := data.(*entities.ProductRequest)

	dataQuery := fmt.Sprintf("%s %s %s (%s) %s (?,?,?)", INSERT, INTO, entity, dataFields, VALUES)
	stmt, err := s.conexao.Prepare(dataQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(&product.Name, &product.Price, &product.Count)
	if err != nil {
		return err
	}

	if rowsAffected, err := result.RowsAffected(); err != nil && rowsAffected == 0 {
		return err
	}
	return nil
}

func (s SQLite3Repository) GetEntityAll(entity string, fields []string) interface{} {
	dataFields := strings.Join(fields, ",")
	dataQuery := fmt.Sprintf("%s %s %s %s", SELECT, dataFields, FROM, entity)
	result, err := s.conexao.Query(dataQuery)
	if err != nil {
		return nil
	}
	s.resultSet = result
	return s.resultSet
}

func (s SQLite3Repository) GetEntityByID(entity string, fields []string, data interface{}) interface{} {
	dataFields := strings.Join(fields, ",")
	product := data.(*entities.ProductRequest)
	query := fmt.Sprintf("%s %s %s %s %s %s=?", SELECT, dataFields, FROM, entity, WHERE, fields[0])
	s.resultSet = s.conexao.QueryRow(query, product.ID)
	return s.resultSet
}

func (s *SQLite3Repository) UpdateEntity(tableName string, fields []string, oldData, data interface{}) error {
	product := data.(*entities.ProductRequest)
	oldProduct := oldData.(*entities.ProductRequest)
	if oldProduct.Count+product.Count >= 0 {
		querySql := fmt.Sprintf("%s %s %s %s=? %s %s=?", UPDATE, tableName, SET, fields[1], WHERE, fields[0])
		oldProduct.Count = oldProduct.Count + product.Count
		result, err := s.conexao.Exec(querySql, &oldProduct.Count, &oldProduct.ID)
		if err != nil {
			return err
		}
		if rowsAffected, err := result.RowsAffected(); err != nil && rowsAffected == 0 {
			return err
		}
	}
	return nil
}

func (s SQLite3Repository) DeleteEntity(entity string, fields []string, data interface{}) error {
	product := data.(*entities.ProductRequest)
	query := fmt.Sprintf("%s %s %s %s %s=?", DELETE, FROM, entity, WHERE, fields[0])
	result, err := s.conexao.Exec(query, product.ID)
	if err != nil {
		return err
	}
	if rowsAffected, err := result.RowsAffected(); err != nil && rowsAffected == 0 {
		return err
	}
	return nil
}
