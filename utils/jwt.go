package utils

import (
	"gin_admin_api/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims JWT结构
type Claims struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}

// GenerateToken JWT生成token
func GenerateToken(id int) (string, error) {
	claims := Claims{
		Id: id,
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
