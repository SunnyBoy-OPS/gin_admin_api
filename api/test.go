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
// @Tags 用户登录接口
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
	LoginQuery(c)
}
