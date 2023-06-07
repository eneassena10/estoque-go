package main

import (
	"database/sql"

	config "github.com/eneassena10/estoque-go/internal/configuracao"
	productController "github.com/eneassena10/estoque-go/internal/domain/product/controllers"
	userController "github.com/eneassena10/estoque-go/internal/domain/user/controllers"
	service_user "github.com/eneassena10/estoque-go/internal/domain/user/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// create instance
	route := gin.Default()

	database := &sql.DB{}

	service := service_user.NewServiceUser(database)
	app := config.NewApp(
		productController.NewControllers(database),
		userController.NewUserController(service),
	)
	app.InitApp(route)

	// start app
	route.Run(":8080")
}
