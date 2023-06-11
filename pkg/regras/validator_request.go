package regras

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strings"

	// model_products "web-service-gin/internal/products/model"

	"github.com/eneassena10/estoque-go/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RequestError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ResponseError struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "numeric":
		return "This field only accepts numbers"
	case "alphanum":
		return "This field only accepts alphanumeric"
	}
	return ""
}

func ValidateErrorInRequest(context *gin.Context, data interface{}) bool {
	var out []RequestError
	err := context.BindJSON(&data)
	if err != nil {
		var validatorError validator.ValidationErrors
		var jsonError *json.UnmarshalTypeError
		var jsonFieldError *json.UnmarshalFieldError
		switch {
		case errors.As(err, &jsonError):

			strin := strings.Split(jsonError.Error(), ":")[1]
			req := RequestError{jsonError.Field, strin}
			context.JSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, req))

		case errors.As(err, &validatorError):
			out = make([]RequestError, len(validatorError))

			typeData := reflect.TypeOf(data).Elem()
			for i, fe := range validatorError {

				field, _ := typeData.FieldByName(fe.Field())
				out[i] = RequestError{field.Tag.Get("json"), msgForTag(fe.Tag())}
			}
			context.JSON(http.StatusUnprocessableEntity,
				web.NewResponse(http.StatusUnprocessableEntity, out))

		case errors.As(err, &jsonFieldError):
			strin := strings.Split(jsonError.Error(), ":")[1]
			req := RequestError{jsonError.Field, strin}
			context.JSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, req))
		default:
			context.JSON(http.StatusUnprocessableEntity,
				web.DecodeError(http.StatusUnprocessableEntity, err.Error()))
		}
		return true
	}
	return false
}
