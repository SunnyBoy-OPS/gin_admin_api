package main

// 获取前端api的简单示例
import (
	"gin_admin_api/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func AdminVue() {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/sysAdmin/login", func(c *gin.Context) {
			var req LoginReq
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
				return
			}

			// TODO: 校验账号密码
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "ok",
				"data": gin.H{"token": "demo-token"},
			})
		})
	}

	if err := r.Run(":5001"); err != nil {
		global.Log.Error(err)
	}
}
