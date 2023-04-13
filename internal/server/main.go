package main

import (
	"github.com/eneassena10/estoque-go/internal/configuracao"
	"github.com/gin-gonic/gin"
)

func main() {
	// create instance
	r := gin.Default()

	// configura as rotas com os controllers
	configuracao.InitApp(r)

	// start app
	r.Run(":8080")
}
