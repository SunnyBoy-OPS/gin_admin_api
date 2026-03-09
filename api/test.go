// Package api 测试相关接口
// @author Flank
package api

import (
	"gin_admin_api/result"

	"github.com/gin-gonic/gin"
)

// Success 成功测试
// @Summary 成功测试接口
// @Tags 测试相关接口
// @Produce JSON
// @Description 成功测试接口
// @Success 200 {object} result.Result
// @router /api/success [get]
func Success(c *gin.Context) {
	result.Success(c, 200)
}

// Failed 失败测试
// @Summary 失败测试接口
// @Tags 测试相关接口
// @Produce JSON
// @Description 失败测试接口
// @Failed 200 {object} result.Result
// @router /api/failed [get]
func Failed(c *gin.Context) {
	result.Failed(c, int(result.ApiCode.Fail), result.ApiCode.GetMessage(result.ApiCode.Fail))
}

// Failed 失败测试
// @Summary 失败测试接口
// @Tags 测试相关接口
// @Produce JSON
// @Description 失败测试接口
// @Failed 200 {object} result.Result
// @router /api/failed [get]
