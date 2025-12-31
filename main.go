// 启动程序
// @author Flank

package main

import (
	"fmt"
	"gin_admin_api/config"
	"gin_admin_api/core"
	"gin_admin_api/global"
	router2 "gin_admin_api/router"
)

// @title Admin 管理系统
// @version 1.0
// @description Admin后台管理系统
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	global.Log = core.InitLogger()
	//global.Log.Infof("mysql配置: %s", config.Config.Mysql)
	//初始化mysql
	core.MysqlInit()

	//初始化启动路由
	router := router2.InitRouter()
	address := fmt.Sprintf("%s:%s", config.Config.System.Host, config.Config.System.Port)
	global.Log.Infof("系统启动成功，允许在：%s", address)
	err := router.Run(address)
	if err != nil {
		return
	}
	// 初始化logger
	//global.Log = core.InitLogger()
	//global.Log.Warnln("go的日志")
	//global.Log.Error("go的日志")
	//global.Log.Infof("go的日志")
}
