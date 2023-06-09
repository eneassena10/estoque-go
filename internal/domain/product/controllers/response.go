package controllers

//go:generate mockgen -source=./response.go -destination=./mocks/mock_response.go -package=mockgen
type Response struct {
	Code int
	Data interface{}
}

type ProductRequestBody struct {
	ID         int     `json:"id,omitempty"`
	Name       string  `json:"name,omitempty"`
	Price      float64 `json:"price,omitempty"`
	Quantidade int     `json:"quantidade,omitempty"`
}
