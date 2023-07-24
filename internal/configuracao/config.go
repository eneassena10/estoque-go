package configuracao

import (
	"github.com/eneassena10/estoque-go/internal/domain/product/controllers"
	repositoryProduct "github.com/eneassena10/estoque-go/internal/domain/product/repository"
	sqlite3 "github.com/eneassena10/estoque-go/internal/domain/product/repository/database"
	serviceProduct "github.com/eneassena10/estoque-go/internal/domain/product/service"
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

	// dependencias do service de product
	repositoryProduct := repositoryProduct.NewProductRepository(operationsSqlite3)
	serviceProduct := serviceProduct.NewProductService(repositoryProduct)
	repoProduct := controllers.NewControllers(serviceProduct)

	// mappper services
	services := map[HandleNameType]interface{}{
		HandleProduct: repoProduct,
	}

	app := NewApp(services)
	return app
}
