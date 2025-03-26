package router

import (
	"fuck-the-world/config"
	"fuck-the-world/internal/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			项目文档
//	@version		1.0
//	@description	这是一个简单的记账项目

//	@contact.name	herzorf
//	@contact.url	https://github.com/herzorf
//	@contact.email	herzorf@icloud.com

//	@host		localhost:8080
//	@BasePath	/

func New() *gin.Engine {
	config.LoadConfigYaml()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 允许前端的地址
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // 允许携带 Cookie
	}))

	{
		v1 := r.Group("/api/v1")

		v1.POST("/sendEmail", controller.SendEmail)
		v1.POST("/login", controller.Login)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func RunServer() {
	r := New()
	err := r.Run("0.0.0.0:8888")
	if err != nil {
		panic(err)
	}
}
