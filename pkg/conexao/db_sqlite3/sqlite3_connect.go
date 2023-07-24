package dbsqlite3

import (
	"database/sql"
	"log"
)

type IDataBase interface {
	DBConnect() *sql.DB
}

//go:generate mockgen -source=./entites.go -destination=./../../../test/mockgen/entites_mock.go -package=mockgen
type IDataBaseOperation interface {
	GetEntityAll(entity string, fields []string) interface{}
	GetEntityByID(entity string, fields []string, data interface{}) interface{}
	CreateEntity(entity string, fields []string, data interface{}) error
	UpdateEntity(entity string, oldData interface{}, data interface{}) error
	DeleteEntity(entity string, fields []string, data interface{}) error
}

const PATH_DB = "../pkg/conexao/db_sqlite3/data/service.db"

func DBConnect() *sql.DB {
	database, err := sql.Open("sqlite3", PATH_DB)
	if err != nil {
		log.Fatal(err)
	}

	if err := database.Ping(); err != nil {
		log.Fatal(err)
	}
	return database
}
