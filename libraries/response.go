package libraries

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseSuccess struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
type ResponseError struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Errors interface{} `json:"errors"`
}

func StatusOk(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseSuccess{
		Code:   200,
		Status: "OK",
		Data:   data,
	})
}
func StatusCreated(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, &ResponseSuccess{
		Code:   201,
		Status: "Created",
		Data:   data,
	})
}
func StatusBadRequest(c *gin.Context, errors interface{}) {
	c.JSON(http.StatusBadRequest, &ResponseError{
		Code:   400,
		Status: "Bad Request",
		Errors: errors,
	})
}
func StatusNotFound(c *gin.Context, errors interface{}) {
	c.JSON(http.StatusNotFound, &ResponseError{
		Code:   404,
		Status: "Not Found",
		Errors: errors,
	})
}
func StatusInternalServerError(c *gin.Context, errors interface{}) {
	c.JSON(http.StatusInternalServerError, &ResponseError{
		Code:   500,
		Status: "Internal Server Error",
		Errors: errors,
	})
}
func StatusNoContent(c *gin.Context, data interface{}) {
	c.JSON(http.StatusNoContent, &ResponseSuccess{
		Code:   204,
		Status: "No Content",
		Data:   data,
	})
}
