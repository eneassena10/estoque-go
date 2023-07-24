package main

import (
	"github.com/eneassena10/estoque-go/internal/configuracao"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	route := gin.New()

	route.Use(gin.Logger())

	configuracao.Start().Routers(route)

	route.Run(":8080")
}
