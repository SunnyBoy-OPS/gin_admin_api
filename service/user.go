package service

import (
	"gin_admin_api/utils"
)

// LoginToken 利用用户Id来获取token
func LoginToken(id int) (string, error) {

	// 生成token
	token, err := utils.GenerateToken(id)
	if err != nil {
		return "", err
	}
	return token, err
}
