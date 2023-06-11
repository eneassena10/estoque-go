package sqlite3_repository

import (
	"database/sql"
)

const PATH_DB = "./loja.db"

func DBConnect() *sql.DB {
	db, err := sql.Open("sqlite3", PATH_DB)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}
