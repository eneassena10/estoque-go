package controllers

type Product struct {
	ID         int     `json:"id,omitempty"`
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
	{
		ID:         3,
		Name:       "Mouse",
		Price:      55,
		Quantidade: 60,
	},
}

type Response struct {
	Code int
	Data interface{}
}

var mapMessageHttp = map[int]string{
	200: "Sucess",
	404: "Not Found",
}
