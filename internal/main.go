package main

import (
	"github.com/eneassena10/estoque-go/internal/auth"
	config "github.com/eneassena10/estoque-go/internal/configuracao"
	"github.com/eneassena10/estoque-go/internal/controllers"
	"github.com/eneassena10/estoque-go/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {
	// create instance
	route := gin.Default()

	fileStore := store.NewFileStore(controllers.ProductPathName)
	app := config.NewApp(
		fileStore,
		controllers.NewControllers(fileStore),
		auth.NewUserController(),
	)
	app.InitApp(route)

	// start app
	route.Run(":8080")
}
