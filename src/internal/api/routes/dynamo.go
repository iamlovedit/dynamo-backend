package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iamlovedit/dynamo-backend/internal/api/handlers"
)

func RegisterDynamoRoutes(router *gin.RouterGroup, handler *handlers.DynamoHandler) {
	dynamoGroup := router.Group("/dynamo")
	dynamoGroup.GET("/:id", handler.GetById)
}
