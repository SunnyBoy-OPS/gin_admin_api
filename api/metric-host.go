package api

import (
	"gin_admin_api/model"
	"gin_admin_api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TotalMetric 通过 SSH 采集主机配置（CPU核数、内存总量、磁盘总量）
func TotalMetric(c *gin.Context) {
	var req model.ConnectTestHost
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "参数错误",
			"err":  err.Error(),
		})
		return
	}

	metric, err := service.ConnectTestMetic(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "采集失败",
			"err":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "采集成功",
		"data": metric,
	})
}
