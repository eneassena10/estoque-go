package sqlite3

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/eneassena10/estoque-go/internal/domain/product/entities"
	dbsqlite3 "github.com/eneassena10/estoque-go/pkg/conexao/db_sqlite3"
)

type SQLite3 struct {
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
	UPDATE = "UPDADE"
	TABLE  = "TABLE"
	SET    = "SET"
)

func NewSQLite3(con *sql.DB) dbsqlite3.IDataBaseOperation {
	return &SQLite3{
		conexao: con,
	}
}

func (s SQLite3) CreateEntity(entity string, fields []string, data interface{}) error {
	dataFields := strings.Join(fields, ",")
	product := data.(*entities.ProductRequest)
	dataQuery := fmt.Sprintf("%s %s %s (%s) %s (?,?,?)", INSERT, INTO, entity, dataFields, VALUES)
	fmt.Println(dataQuery, dataFields)
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

func (s SQLite3) GetEntityAll(entity string, fields []string) interface{} {
	//  entity, fields, estrutura
	dataFields := strings.Join(fields, ",")
	dataQuery := fmt.Sprintf("%s %s %s %s", SELECT, dataFields, FROM, entity)
	result, err := s.conexao.Query(dataQuery)
	log.Println(result)
	if err != nil {
		return nil
	}
	s.resultSet = result
	return s.resultSet
}

func (s SQLite3) GetEntityByID(entity string, fields []string, data interface{}) interface{} {
	dataFields := strings.Join(fields, ",")
	product := data.(*entities.ProductRequest)
	query := fmt.Sprintf("%s %s %s %s %s %s=?", SELECT, dataFields, FROM, entity, WHERE, fields[0])
	s.resultSet = s.conexao.QueryRow(query, product.ID)
	return s.resultSet
}

func (s SQLite3) UpdateEntity(entity string, oldData interface{}, data interface{}) error {
	return nil
}

func (s SQLite3) DeleteEntity(entity string, fields []string, data interface{}) error {
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
