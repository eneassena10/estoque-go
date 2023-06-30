package configuracao

import (
	productController "github.com/eneassena10/estoque-go/internal/domain/product/controllers"
	sqlite3_repository "github.com/eneassena10/estoque-go/internal/domain/product/repository"
	sqlite3 "github.com/eneassena10/estoque-go/internal/domain/product/repository/database"
	"github.com/eneassena10/estoque-go/internal/domain/product/service"
	dbsqlite3 "github.com/eneassena10/estoque-go/pkg/conexao/db_sqlite3"
)

func Start() *App {
	// criar uma conexao
	conexaoDB := dbsqlite3.DBConnect()
	// criar um instancia de operações
	operationsSqlite3 := sqlite3.NewSQLite3(conexaoDB)
	// criar um repository
	productRepository := sqlite3_repository.NewProductRepository(operationsSqlite3)
	// criar um service
	productService := service.NewProductService(productRepository)
	// criar um controller
	productControllers := productController.NewControllers(productService)

	app := NewApp(productControllers)
	return app
}
