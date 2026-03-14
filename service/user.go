package service

import (
	"gin_admin_api/config"
	"gin_admin_api/dao"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims JWT结构
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

// LoginToken 登录逻辑,获取token
func LoginToken(username, password string) (string, error) {
	// 查询用户密码
	_, err := dao.GetUser(username, password)
	if err != nil {
		return "", err
	}

	// 生成token
	token, err := GenerateToken(username, password)
	if err != nil {
		return "", err
	}
	return token, err
}

// GenerateToken JWT生成token
func GenerateToken(username, password string) (string, error) {
	claims := Claims{
		Username: username,
		Password: password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JwtSecret)
}

// ParseToken 解析token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return config.JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(*Claims), nil
}
