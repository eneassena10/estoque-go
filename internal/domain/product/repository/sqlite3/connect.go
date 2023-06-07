package sqlite3

import (
	"database/sql"
	"log"
)

const PATH_DB = "./loja.db"

func DBConnect() *sql.DB {
	db, err := sql.Open("sqlite3", PATH_DB)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
