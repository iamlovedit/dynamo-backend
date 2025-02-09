package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpDataResponse[T any] struct {
	Code    int
	Message string
	Data    T
}

func SucceedData[T any](c *gin.Context, code int, message string, data T) {
	httpResponse := HttpDataResponse[T]{code, message, data}
	c.JSON(http.StatusOK, httpResponse)
}

func Succeed(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": "成功",
		"data":    data,
	})
}

func Failed(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
	})
}
