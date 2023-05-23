package main

import (
	config "github.com/eneassena10/estoque-go/internal/configuracao"
	productController "github.com/eneassena10/estoque-go/internal/product/controllers"
	"github.com/eneassena10/estoque-go/internal/product/domain"
	userController "github.com/eneassena10/estoque-go/internal/user/controllers"
	service_user "github.com/eneassena10/estoque-go/internal/user/service"
	"github.com/eneassena10/estoque-go/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {
	// create instance
	route := gin.Default()

	fileStore := store.NewFileStore(domain.ProductPathName)
	service := service_user.NewServiceUser(fileStore)
	app := config.NewApp(
		fileStore,
		productController.NewControllers(fileStore),
		userController.NewUserController(service),
	)
	app.InitApp(route)

	// start app
	route.Run(":8080")
}
