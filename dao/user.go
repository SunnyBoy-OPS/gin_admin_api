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

// GetUserId 生成token需要根据用户名Id
//func GetUserId(id int) (int, error) {
//	var userList []model.User
//	err := core.Db.Where("UserName = ? and UserPassword = ?", id).Find(&userList).Error
//	return userList[0].Id, err
//}
