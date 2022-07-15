package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JsonSuccess return json success
func JsonSuccess(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Success())
}

// JsonSuccessMessage return json success with message
func JsonSuccessMessage(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, SuccessMessage(message))
}

// JsonError return json error
func JsonError(ctx *gin.Context, code int) {
	ctx.JSON(http.StatusOK, Error(code))
}

// JsonErrorMessage return json error with message
func JsonErrorMessage(ctx *gin.Context, code int, error string) {
	ctx.JSON(http.StatusOK, ErrorMessage(code, error))
}

// JsonFail return json fail
func JsonFail(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Fail())
}

// JsonFailMessage return json fail with message
func JsonFailMessage(ctx *gin.Context, error string) {
	ctx.JSON(http.StatusOK, FailMessage(error))
}

// JsonData return json data
func JsonData(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, Data(data))
}

// JsonUnauthorized return json unauthorized
func JsonUnauthorized(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Error(http.StatusUnauthorized))
}

// JsonUnauthorizedMessage return json unauthorized with message
func JsonUnauthorizedMessage(ctx *gin.Context, error string) {
	ctx.JSON(http.StatusOK, ErrorMessage(http.StatusUnauthorized, error))
}
