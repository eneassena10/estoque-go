package controllers

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
