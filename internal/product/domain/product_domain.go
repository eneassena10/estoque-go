package domain

type Product struct {
	ID         int
	Name       string
	Price      float64
	Quantidade int
}

type ProductRequest struct {
	ID         int     `json:"id,omitempty"`
	Name       string  `json:"name" binding:"requered"`
	Price      float64 `json:"price" binding:"requered"`
	Quantidade int     `json:"quantidade" binding:"requered"`
}

var ProductPathName = "./products.json"
