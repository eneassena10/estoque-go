package main

import (
	"github.com/eneassena10/estoque-go/internal/configuracao"
	"github.com/gin-gonic/gin"
)

/*
	adicionar routes : POST, PATH: ok
	mover para um arquivo separado a lógica do controller e declaração de rotas
	pesquisar sobre solid e implementar no projeto um exemplo
	clean arquitecture
	altera os código http
	paradigma de programação kiss
*/

func main() {
	// create instance
	r := gin.Default()

	// configura as rotas com os controllers
	configuracao.InitApp(r)

	// start app
	r.Run(":8080")
}
