package controllers_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	productController "github.com/eneassena10/estoque-go/internal/domain/product/controllers"
	userController "github.com/eneassena10/estoque-go/internal/domain/user/controllers"

	config "github.com/eneassena10/estoque-go/internal/configuracao"
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
	db := &sql.DB{}
	productController := productController.NewControllers(db)
	userController := userController.NewUserController(db)
	app := config.NewApp(productController, userController)
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
