package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamlovedit/dynamo-backend/internal/models"
	"github.com/iamlovedit/dynamo-backend/internal/services"
)

type DynamoHandler struct {
	dynamoService services.IDynamoService
}

func NewDynamoHandler(dynamoService services.IDynamoService) *DynamoHandler {
	return &DynamoHandler{
		dynamoService: dynamoService,
	}
}

func (h *DynamoHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusOK, models.ErrorMessage{Code: 500, Message: "id不能为空"})
		return
	}
	dynamo, err := h.dynamoService.FindDynamo(id)
	if err != nil {
		ctx.JSON(http.StatusOK, models.ErrorMessage{Code: 404, Message: "节点包不存在"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": dynamo,
	})
}
