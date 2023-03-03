package main

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID         int64
	Name       string
	Price      float64
	Quantidade int
}

var products []Product = []Product{
	{
		ID:         1,
		Name:       "Teclado",
		Price:      30,
		Quantidade: 2,
	},
	{
		ID:         2,
		Name:       "Teclado 2",
		Price:      35,
		Quantidade: 20,
	},
}

type Response struct {
	Code int
	Data interface{}
}

func main() {
	r := gin.Default()
	r.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(200, Response{200, products})
	})

	r.GET("/products/:id", func(ctx *gin.Context) {
		param := ctx.Param("id")
		var paramID int64

		paramID, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			ctx.JSON(500, Response{500, err.Error()})
			return
		}

		product := getProductByProductID(paramID)

		if product != nil {
			ctx.JSON(200, Response{200, product})
			return
		}

		messageNotFound := errors.New("not found the product")

		ctx.JSON(404, Response{404, messageNotFound.Error()})
	})
	r.Run(":8080")
}

func getProductByProductID(productID int64) *Product {
	for _, p := range products {
		if p.ID == productID {
			return &p
		}
	}
	return nil
}
