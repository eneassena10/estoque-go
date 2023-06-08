package controllers

//go:generate mockgen -source=./response.go -destination=./mocks/mock_response.go -package=mockgen
type Response struct {
	Code int
	Data interface{}
}
