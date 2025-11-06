package main

import (
	"log"
	"notification/config"
	"notification/handler"
	"notification/service"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "notification/docs"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 初始化服务
	notificationService := service.NewNotificationService(cfg)

	// 初始化路由
	r := gin.Default()

	// 注册swagger路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 注册处理器
	notificationHandler := handler.NewNotificationHandler(notificationService)

	api := r.Group("/api/v1")
	{
		api.POST("/notify/email", notificationHandler.SendEmail)
		api.POST("/notify/sms", notificationHandler.SendSMS)
		api.POST("/notify/batch", notificationHandler.SendBatch)

	}

	// 启动服务
	port := cfg.Server.Port
	if port == "" {
		port = "9000"
	}

	log.Printf("通知系统启动在端口: %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
