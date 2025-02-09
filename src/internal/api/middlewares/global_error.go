package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// 检查是否有错误需要处理
		if len(c.Errors) > 0 {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "服务异常",
			})

		}
	}
}
