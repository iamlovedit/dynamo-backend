package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iamlovedit/dynamo-backend/internal/api/handlers"
	"github.com/iamlovedit/dynamo-backend/internal/api/middlewares"
	"github.com/iamlovedit/dynamo-backend/internal/api/routes"
	"github.com/iamlovedit/dynamo-backend/internal/repository"
	"github.com/iamlovedit/dynamo-backend/internal/services"
)

func main() {
	// 初始化路由
	router := gin.New()

	// 中间件
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	router.Use(middlewares.ErrorMiddleware())

	// 依赖注入
	dynamoRepo := repository.NewMongoDynamoRepository()
	dynamoService := services.NewDynamoService(dynamoRepo)
	dynamoHandler := handlers.NewDynamoHandler(dynamoService)

	// API 路由组
	api := router.Group("/api/v1")
	routes.RegisterDynamoRoutes(api, dynamoHandler)

	// 启动服务器
	router.Run(":8080")
}
