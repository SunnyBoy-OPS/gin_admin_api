package middleware

import (
	"gin_admin_api/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuth JWT验证
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "未携带token",
			})
			c.Abort()
		}

		// 去掉Bearer
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "token无效",
			})
			return
		}
		c.Set("id", claims.Id)

		c.Next()
	}
}
