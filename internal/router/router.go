package router

import (
	_ "kumiko/docs"
	"kumiko/internal/controller"
	"kumiko/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title kumiko API
// @version 1.0
// @description Gin + GORM 项目模板
// @host localhost:8080
// @BasePath /

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/login", controller.Login)

	auth := r.Group("/")
	auth.Use(middleware.JWT())
	auth.GET("/user/:id", controller.GetUser)
	auth.POST("/user", controller.CreateUser)
	auth.GET("/test", controller.Test)
	auth.GET("/user", controller.GetUserList)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
