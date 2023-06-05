package main

import (
	"database/sql"

	config "github.com/eneassena10/estoque-go/internal/configuracao"
	productController "github.com/eneassena10/estoque-go/internal/domain/product/controllers"
	"github.com/eneassena10/estoque-go/internal/domain/product/entities"
	userController "github.com/eneassena10/estoque-go/internal/domain/user/controllers"
	service_user "github.com/eneassena10/estoque-go/internal/domain/user/service"
	"github.com/eneassena10/estoque-go/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {
	// create instance
	route := gin.Default()

	fileStore := store.NewFileStore(entities.ProductPathName)
	service := service_user.NewServiceUser(fileStore)
	app := config.NewApp(
		fileStore,
		productController.NewControllers(fileStore, &sql.DB{}),
		userController.NewUserController(service),
	)
	app.InitApp(route)

	// start app
	route.Run(":8080")
}
