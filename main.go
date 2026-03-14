// 启动程序
// @author Flank
// swagger： http://127.0.0.1:5001/swagger/index.html
package main

import (
	"fmt"
	"gin_admin_api/config"
	"gin_admin_api/core"
	_ "gin_admin_api/docs"
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
	err := core.MysqlInit()
	if err != nil {
		global.Log.Info(err)
		return
	}
	//初始化启动路由
	router := router2.InitRouter()
	address := fmt.Sprintf("%s:%s", config.Config.System.Host, config.Config.System.Port)
	global.Log.Infof("系统启动成功，允许在：%s", address)
	err1 := router.Run(address)
	if err1 != nil {
		global.Log.Info(err1)
		return
	}

	// 初始化logger
	//global.Log = core.InitLogger()
	//global.Log.Warnln("go的日志")
	//global.Log.Error("go的日志")
	//global.Log.Infof("go的日志")
}
