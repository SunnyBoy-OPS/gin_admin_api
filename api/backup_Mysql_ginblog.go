package api

import (
	"fmt"
	"gin_admin_api/config"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh"
)

func BackupMySQL(c *gin.Context) {
	outDir := "./SQL/backup"
	//outDir := `G:\staudy\project\gin_vue3博客\代码\gin_admin_api\SQL`
	var dbcfg = config.Config.Mysql
	var sshcfg = config.Config.SshConfig

	// 准备本地输出目录
	if err := os.MkdirAll(outDir, 0755); err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "准备本地输出目录失败",
			"err":  err.Error(),
			"data": outDir,
		})
		//fmt.Println(err)
	}

	fileName := fmt.Sprintf("%s_%s.sql", dbcfg.Db, time.Now().Format("20060102_150405"))
	localPath := filepath.Join(outDir, fileName)

	// 建立ssh连接
	clientConfig := &ssh.ClientConfig{
		User:            sshcfg.User,
		Auth:            []ssh.AuthMethod{ssh.Password(sshcfg.Password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         30 * time.Second,
	}
	addr := fmt.Sprintf("%s:%d", sshcfg.Host, sshcfg.Port)
	client, err := ssh.Dial("tcp", addr, clientConfig)

	//fmt.Printf("ssh 连接失败: %s", err)
	defer func(client *ssh.Client) {
		err := client.Close()
		if err != nil {
			return
		}
	}(client)

	// 创建远程会话
	session, err := client.NewSession()
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "ssh 连接失败",
			"err":  err.Error(),
		})
		//fmt.Printf("ssh 远程会话创建失败: %s", err)
	}
	defer func(session *ssh.Session) {
		err := session.Close()
		if err != nil {
			return
		}
	}(session)

	// 打开本地文件
	outfile, err := os.Create(localPath)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "ssh 连接失败",
			"err":  err.Error(),
		})
		return
		//fmt.Println(err)
	}
	defer func(outfile *os.File) {
		err := outfile.Close()
		if err != nil {
			return
		}
	}(outfile)

	//远程命令的 stdout 直接写入本地 SQL 文件
	stdout, err := session.StdoutPipe()
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "获取远程 stdout 失败",
			"err":  err.Error(),
		})
		return
		//fmt.Printf("获取远程 stdout 失败: %s", err)
	}
	session.Stderr = os.Stderr

	// 组装远程 mysqldump 命令（在 Linux 上执行）
	cmd := fmt.Sprintf(
		"MYSQL_PWD='%s' mysqldump -h %s -P %d -u %s  --databases %s --single-transaction --set-gtid-purged=OFF", dbcfg.Password, dbcfg.Host, dbcfg.Port, dbcfg.Username, dbcfg.Db,
	)

	//启动命令并接收输出
	if err := session.Start(cmd); err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "启动远程备份失败",
			"err":  err.Error(),
		})
		return
		//fmt.Printf("启动远程备份失败: %s", err)
	}
	if _, err := io.Copy(outfile, stdout); err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "写入本地备份文件失败",
			"err":  err.Error(),
		})
		return
		//fmt.Printf("写入本地备份文件失败: %s", err)
	}
	if err := session.Wait(); err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "远程备份执行失败",
			"err":  err.Error(),
		})
		return
		//fmt.Printf("远程备份执行失败: %s", err)
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "备份成功",
		"data": localPath,
	})
}
