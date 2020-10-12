package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewError(error error) ErrorResponse {
	e := ErrorResponse{
		Error: "",
	}
	if error != nil {
		e.Error = error.Error()
	}
	return e
}

func (err ErrorResponse) SendIfError(context *gin.Context) bool {
	if err.Error == "" {
		return true
	}
	context.JSON(http.StatusInternalServerError, err.Error)
	return false
}
