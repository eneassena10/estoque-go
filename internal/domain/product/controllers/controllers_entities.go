package controllers

type Response struct {
	Code int
	Data interface{}
}
type ProductRequestBody struct {
	ID    int     `json:"id,omitempty" binding:"required"`
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
	Count int     `json:"count,omitempty"`
}
