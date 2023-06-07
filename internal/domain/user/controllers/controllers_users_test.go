package controllers_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	userController "github.com/eneassena10/estoque-go/internal/domain/user/controllers"
	service_user "github.com/eneassena10/estoque-go/internal/domain/user/service"

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
	service := service_user.NewServiceUser(&sql.DB{})
	userController := userController.NewUserController(service)
	app := config.NewApp(nil, userController)
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

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func TestControllers_Logar(t *testing.T) {
	t.Run("Logar", func(t *testing.T) {
		uri := "user/login"
		data := struct {
			NickName string
			Password string
		}{}
		data.NickName = "test"
		data.Password = "testsenha"
		body, _ := json.Marshal(&data)

		rr := CreateServer(http.MethodPost, uri, string(body))
		responseContent := response{}
		_ = json.Unmarshal(rr.Body.Bytes(), &responseContent)

		assert.EqualValues(t, responseContent.Message, "login realizado")
	})
}
