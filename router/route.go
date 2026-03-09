// Package router 初始化路由及注册
// @author Flank
package router

import (
	"gin_admin_api/api"
	"gin_admin_api/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	// 设置启动模式
	gin.SetMode(config.Config.System.Env)
	router := gin.New()

	//宕机恢复
	router.Use(gin.Recovery())
	//register注册
	register(router)
	return router
}

// register 路由接口
func register(router *gin.Engine) {
	// todo 后续接口url
	// 给你的 Gin 项目注册一个路由，用来访问 Swagger API 文档页面
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/api/success", api.SuccessApi)
	router.GET("/api/failed", api.FailedApi)

	// 登录接口
	router.GET("/api/user", api.LoginQuery)
	router.POST("/api/login", api.LoginQuery)

	// 备份数据库接口
	router.POST("/api/backupMysql", api.BackupMySQL)
}
