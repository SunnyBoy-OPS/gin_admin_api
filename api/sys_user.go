package api

import (
	"gin_admin_api/core"
	"gin_admin_api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueryUser(username string, password string) ([]model.User, error) {
	var userList []model.User
	err := core.Db.Where("UserName = ? and UserPassword = ?", username, password).Find(&userList).Error
	return userList, err
}
func Test(c *gin.Context) {
	var users []model.User

	// 获取url上的username，password参数值
	username := c.Query("username")
	password := c.Query("password")

	// 判断拿到的username，password是否是空值，是直接返回前端
	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "账号密码不能为空！",
		})
	}
	// 获取到的值传给QueryUser 去数据库查询
	users, err := QueryUser(username, password)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"code":  http.StatusServiceUnavailable,
			"msg":   "服务器或者数据库查询失败",
			"error": err.Error(),
		})
		return
	} else if len(users) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "账号密码错误",
			"data": users,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "查询成功",
		"count": len(users),
		"data":  users,
	})
	return
}
