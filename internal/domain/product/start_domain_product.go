package product

import (
	"github.com/eneassena10/estoque-go/internal/domain/product/controllers"
	repositoryUser "github.com/eneassena10/estoque-go/internal/domain/product/repository"
	serviceUser "github.com/eneassena10/estoque-go/internal/domain/product/service"
	dbsqlite3 "github.com/eneassena10/estoque-go/pkg/conexao/db_sqlite3"
)

type serviceDomainProduct struct {
	Handler *controllers.ProductControllers
}

func StartDomain(database dbsqlite3.IDataBaseOperation) *serviceDomainProduct {
	repositoryUser := repositoryUser.NewProductRepository(database)
	serviceUser := serviceUser.NewProductService(repositoryUser)
	repositoryProduct := controllers.NewControllers(serviceUser)
	return &serviceDomainProduct{
		Handler: repositoryProduct,
	}
}
