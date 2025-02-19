package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/iamlovedit/dynamo-backend/internal/models"
	"net/http"
)

func NotFoundMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, models.ErrorMessage{
			Code:    404,
			Message: "接口不存在",
		})
	}
}
