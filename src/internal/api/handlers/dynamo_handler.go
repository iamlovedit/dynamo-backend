package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/iamlovedit/dynamo-backend/internal/services"
)

type DynamoHandler struct {
	dynamoService services.DynamoService
}

func NewDynamoHandler(dynamoService services.DynamoService) *DynamoHandler {
	return &DynamoHandler{
		dynamoService,
	}
}

func (h *DynamoHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	h.dynamoService.FindDynamo(id)
}
