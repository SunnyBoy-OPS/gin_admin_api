// Package api 登录没有token接口
package api

import (
	"gin_admin_api/dao"
	"gin_admin_api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginQuery 处理登录的api
func LoginQuery(c *gin.Context) {
	//c := new(gin.Context)
	var userList []model.User

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
	// 获取到的值传给QueryUser 去数据库查询
	userList, err := dao.GetUser(username, password)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"code":  http.StatusServiceUnavailable,
			"msg":   "服务器或者数据库查询失败",
			"error": err.Error(),
		})
		return
	} else if len(userList) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "账号密码错误",
			"data": userList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "登录成功",
		"count": len(userList),
		"data":  userList,
	})
	return
}
