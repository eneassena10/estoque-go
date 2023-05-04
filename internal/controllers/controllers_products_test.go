package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eneassena10/estoque-go/internal/auth"
	config "github.com/eneassena10/estoque-go/internal/configuracao"
	"github.com/eneassena10/estoque-go/internal/controllers"
	"github.com/eneassena10/estoque-go/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type ResponseContent struct {
	Code  int
	Data  interface{}
	Error string
}

type Response struct {
	Code int
	Data interface{}
}

func CreateServer(method, url, body string) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	productNamePath := "../products.json"
	fileStore := store.NewFileStore(productNamePath)
	handlers := controllers.NewControllers(fileStore)
	userController := auth.NewUserController()
	app := config.NewApp(fileStore, handlers,userController)
	app.InitApp(router)
	req, rr := CreateRequestTest(method, url, body)
	router.ServeHTTP(rr, req)
	return rr
}

func CreateRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func TestControllers_GetProductsAll(t *testing.T) {
	t.Run("GetProductsAll", func(t *testing.T) {
		uri := "/products/list"
		rr := CreateServer(http.MethodGet, uri, "")
		responseContent := Response{}
		_ = json.Unmarshal(rr.Body.Bytes(), &responseContent)
		assert.Len(t, responseContent.Data, 5)
		assert.EqualValues(t, responseContent.Code, 200)
	})
}
