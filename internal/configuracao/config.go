package configuracao

import (
	"github.com/eneassena10/estoque-go/internal/domain/product/controllers"
	repositoryProduct "github.com/eneassena10/estoque-go/internal/domain/product/repository"
	sqlite3 "github.com/eneassena10/estoque-go/internal/domain/product/repository/database"
	serviceProduct "github.com/eneassena10/estoque-go/internal/domain/product/service"
	dbsqlite3 "github.com/eneassena10/estoque-go/pkg/conexao/db_sqlite3"
)

const (
	DomainProduct = "DomainProduct"
	DomaineUser   = "DomainUser"
)

func Start() *App {
	conexaoDB := dbsqlite3.DBConnect()
	operationsSqlite3 := sqlite3.NewSQLite3(conexaoDB)

	repositoryUser := repositoryProduct.NewProductRepository(operationsSqlite3)
	serviceUser := serviceProduct.NewProductService(repositoryUser)
	repoProduct := controllers.NewControllers(serviceUser)

	services := map[string]interface{}{
		DomainProduct: repoProduct,
	}

	app := NewApp(services)
	return app
}
