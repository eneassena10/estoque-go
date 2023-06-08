package main

import (
	"github.com/eneassena10/estoque-go/internal/configuracao"
	productController "github.com/eneassena10/estoque-go/internal/domain/product/controllers"
	sqlite3_repository "github.com/eneassena10/estoque-go/internal/domain/product/repository/sqlite3"
	userController "github.com/eneassena10/estoque-go/internal/domain/user/controllers"
	service_user "github.com/eneassena10/estoque-go/internal/domain/user/service"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// create instance
	route := gin.Default()

	database := sqlite3_repository.DBConnect()

	service := service_user.NewServiceUser(database)
	app := configuracao.NewApp(
		productController.NewControllers(database),
		userController.NewUserController(service),
	)
	app.InitApp(route)

	// start app
	route.Run(":8080")
}
