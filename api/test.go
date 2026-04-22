// Package api 测试相关接口
// @author Flank
package api

import (
	"gin_admin_api/result"

	"github.com/gin-gonic/gin"
)

// SuccessApi 成功测试
// @Summary 成功测试接口
// @Tags 测试相关接口
// @Produce json
// @Description 成功测试接口
// @Success 200 {object} result.Result
// @Router /api/success [get]
func SuccessApi(c *gin.Context) {
	result.Success(c, 200)
}

// FailedApi 失败测试
// @Summary 失败测试接口
// @Tags 测试相关接口
// @Produce json
// @Description 失败测试接口
// @Failure 200 {object} result.Result
// @Router /api/failed [get]
func FailedApi(c *gin.Context) {
	result.Failed(c, int(result.ApiCode.Fail), result.ApiCode.GetMessage(result.ApiCode.Fail))
}

// LoginApi 用户登录
// @Summary 用户登录
// @Description 根据用户名和密码登录
// @Tags 用户登录接口--不生成token
// @Accept json
// @Produce json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {object} result.Result
// @Failure 400 {object} map[string]interface{}
// @Failure 503 {object} map[string]interface{}
// @Router /api/login [post]
func LoginApi(c *gin.Context) {
	// LoginReq 处理登录的api
	//LoginQuery(c)
	LoginQuery(c)
}

// Backup_ginblogApi ginblog数据库备份接口
// @Summary 数据库备份
// @Description 通过SSH在Linux执行mysqldump并保存到本地SQL目录
// @Tags 数据库备份接口
// @Accept json
// @Produce json
// @Success 200 {object} result.Result
// @Failure 400 {object} map[string]interface{}
// @Failure 503 {object} map[string]interface{}
// @Router /api/backupMysql [post]
func Backup_ginblogApi(c *gin.Context) {
	// LoginReq 处理登录的api
	BackupMySQL(c)
}

// LoginTokenApi LoginToken LoginApi 用户登录
// @Summary 用户登录
// @Description 根据用户名和密码登录
// @Tags 用户登录接口--生成token
// @Accept json
// @Produce json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {object} result.Result
// @Failure 400 {object} map[string]interface{}
// @Failure 503 {object} map[string]interface{}
// @Router /api/logintoken [post]
func LoginTokenApi(c *gin.Context) {
	LoginToken(c)
}

// ConnectTestHostApi 测试连接主机（host/ip、端口、账号、密码）
// @Summary 测试连接主机
// @Description 测试连接主机（host/ip、端口、账号、密码）
// @Tags 测试连接主机--（host/ip、端口、账号、密码）
// @Accept json
// @Produce json
// @Param data body model.ConnectTestHost true "通过 SSH 测试连接主机"
// @Success 200 {object} result.Result
// @Failure 400 {object} map[string]interface{}
// @Failure 503 {object} map[string]interface{}
// @Router /api/connect-test-host [post]
func ConnectTestHostApi(c *gin.Context) {
	ConnectTestHost(c)
}

// TotalMetricApi 通过 SSH 采集主机配置（CPU核数、内存总量、磁盘总量）
// @Summary 通过 SSH 采集主机配置（CPU核数、内存总量、磁盘总量）
// @Description 通过 SSH 采集主机配置（CPU核数、内存总量、磁盘总量）
// @Tags 测试连接主机--（host/ip、端口、账号、密码）
// @Accept json
// @Produce json
// @Param data body model.ConnectTestHost true "通过 SSH 采集主机配置"
// @Success 200 {object} result.Result
// @Failure 400 {object} map[string]interface{}
// @Failure 503 {object} map[string]interface{}
// @Router /api/total-metric [post]
func TotalMetricApi(c *gin.Context) {
	TotalMetric(c)
}
