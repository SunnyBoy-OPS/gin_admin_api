// Package api 处理登录接口--登录生成token接口
package api

import (
	"gin_admin_api/dao"
	"gin_admin_api/model"
	"gin_admin_api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginToken 处理登录接口
func LoginToken(c *gin.Context) {
	var userlist []model.User

	// 获取url上的username，password参数值
	username := c.Query("username")
	password := c.Query("password")

	// 判断拿到的username，password是否是空值，是直接返回前端
	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "账号密码不能为空！",
		})
		return
	}

	// url获取到的值传给QueryUser 调用GetUser数据库查询
	userlist, err := dao.GetUser(username, password)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"code":  http.StatusServiceUnavailable,
			"msg":   "服务器或者数据库查询失败!",
			"error": err.Error(),
		})
	} else if len(userlist) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "账号密码错误!",
			"data": userlist,
		})
		return
	}

	//获取到token
	token, err := service.LoginToken(username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"msg":   "账户密码错误！！！",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"data":  userlist,
		"msg":   "登录成功",
		"token": token,
	})
	return
}
