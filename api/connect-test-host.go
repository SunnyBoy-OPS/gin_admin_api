package api

import (
	"fmt"
	"gin_admin_api/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh"
)

// ConnectTestHost 测试连接主机（host/ip、端口、账号、密码）
func ConnectTestHost(c *gin.Context) {
	var req model.ConnectTestHost
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "参数错误",
			"err":  err.Error(),
		})
		return
	}

	clientConfig := &ssh.ClientConfig{
		User:            req.Username,
		Auth:            []ssh.AuthMethod{ssh.Password(req.Password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}

	addr := fmt.Sprintf("%s:%d", req.Host, req.Port)
	client, err := ssh.Dial("tcp", addr, clientConfig)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "ssh 连接失败",
			"err":  err.Error(),
		})
		return
	}
	defer func(client *ssh.Client) {
		err := client.Close()
		if err != nil {
			return
		}
	}(client)

	session, err := client.NewSession()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "创建 ssh 会话失败",
			"err":  err.Error(),
		})
		return
	}
	defer func(session *ssh.Session) {
		err := session.Close()
		if err != nil {
			return
		}
	}(session)

	if err := session.Run("echo ok"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "ssh 命令执行失败",
			"err":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "连接成功",
		"data": gin.H{
			"addr": addr,
			"host": req.Host,
			"port": req.Port,
		},
	})
	return
}
