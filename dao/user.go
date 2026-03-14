package dao

import (
	"gin_admin_api/core"
	"gin_admin_api/model"
)

// GetUser 登录请求参数（Query）根据用户名密码查询用户
func GetUser(username string, password string) ([]model.User, error) {
	var userList []model.User
	err := core.Db.Where("UserName = ? and UserPassword = ?", username, password).Find(&userList).Error
	return userList, err
}
