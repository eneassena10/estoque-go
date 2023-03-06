package main

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID         int64   `json:"id,omitempty"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Quantidade int     `json:"quantidade"`
}

var products []*Product = []*Product{
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
var MessageNotFound error = errors.New("not found the product")

type Response struct {
	Code int
	Data interface{}
}

func main() {
	r := gin.Default()
	r.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(200, Response{200, products})
	})

	r.DELETE("/products/:id", func(ctx *gin.Context) {
		param := ctx.Param("id")
		var paramID int64

		paramID, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			ctx.JSON(500, Response{500, err.Error()})
			return
		}

		prod, err := deleteProductByProductID(paramID)

		products = prod

		if err != nil {
			ctx.JSON(404, Response{404, err.Error()})
			return
		}

		ctx.JSON(204, Response{204, nil})
	})
	r.POST("/products", func(ctx *gin.Context) {
		var product *Product

		if err := ctx.Bind(&product); err != nil {
			ctx.JSON(500, Response{500, err.Error()})
			return
		}

		products = createdProducts(ctx, product)

		ctx.JSON(200, product)
	})
	r.Run(":8080")
}

func deleteProductByProductID(productID int64) ([]*Product, error) {
	for i, p := range products {
		if p.ID == productID {
			if p.Quantidade > 1 {
				p.Quantidade--
				return products, nil
			} else {
				products = append(products[:i], products[i+1:]...)
				return products, nil
			}
		}
	}
	return products, MessageNotFound
}

func createdProducts(ctx *gin.Context, product *Product) []*Product {
	for p := range products {
		if products[p].ID == product.ID {
			products[p].Quantidade += product.Quantidade
			product = products[p]
			return products
		}
	}
	products = append(products, product)
	return products
}
