package configs

import (
	"github.com/eneassena10/estoque-go/internal/domain/product/config"
	sqlite3 "github.com/eneassena10/estoque-go/internal/domain/product/repository/database"
	dbsqlite3 "github.com/eneassena10/estoque-go/pkg/conexao/db_sqlite3"
)

type HandleNameType string

var (
	HandleProduct HandleNameType = "HandleProduct"
	HandleUser    HandleNameType = "HandleUser"
)

func Start() *App {
	return start()
}

func start() *App {
	// conexão com banco de dados
	conexaoDB := dbsqlite3.DBConnect()
	operationsSqlite3 := sqlite3.NewSQLite3(conexaoDB)

	// mappper services
	services := map[HandleNameType]interface{}{
		HandleProduct: config.StartServiceDomainProduct(operationsSqlite3).Handler,
	}

	return NewApp(services)
}
