package controllers

//go:generate mockgen -source=./response.go -destination=./mocks/mock_response.go -package=mockgen
type Response struct {
	Code int
	Data interface{}
}
type ProductRequestBody struct {
	ID         int     `json:"id" binding:"required"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Quantidade int     `json:"quantidade"`
}
