package service

import (
	"fmt"
	"gin_admin_api/model"
	"math"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

func ConnectTestMetic(req model.ConnectTestHost) (*model.CmdInfo, error) {
	// 测试连接主机的时候连接主机
	client := &ssh.ClientConfig{
		User:            req.Username,
		Auth:            []ssh.AuthMethod{ssh.Password(req.Password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         30 * time.Second,
	}

	addr := fmt.Sprintf("%s:%d", req.Host, req.Port)
	conn, err := ssh.Dial("tcp", addr, client)
	if err != nil {
		return nil, err
	}
	defer func(conn *ssh.Client) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	//数据采集
	cpuRaw, err := runSSHCommand(conn, "nproc")
	cpuCores, err := strconv.Atoi(strings.TrimSpace(cpuRaw))
	if err != nil {
		return nil, fmt.Errorf("parse cpu cores failed: %w", err)
	}

	//内存采集
	memRaw, err := runSSHCommand(conn, "awk '/MemTotal/ {print $2}' /proc/meminfo")
	if err != nil {
		return nil, fmt.Errorf("get memory total failed: %w", err)
	}
	memKB, err := strconv.ParseUint(strings.TrimSpace(memRaw), 10, 64)
	//转化GB，保留小数点后面俩位数
	memGB := float64(memKB) / 1024.0 / 1024.0
	memGB = math.Round(memGB*100) / 100
	memGBStr := fmt.Sprintf("%.2f", memGB)
	if err != nil {
		return nil, fmt.Errorf("parse memory total failed: %w", err)
	}

	//磁盘采集
	diskRaw, err := runSSHCommand(conn, "df -kP --total | awk 'END{print $2}'")
	if err != nil {
		return nil, fmt.Errorf("get disk total failed: %w", err)
	}
	//转化GB，保留小数点后面俩位数
	diskKB, err := strconv.ParseUint(strings.TrimSpace(diskRaw), 10, 64)
	diskGB := float64(diskKB) / 1024.0 / 1024.0
	diskGB = math.Round(diskGB*100) / 100
	diskGBStr := fmt.Sprintf("%.2f", diskGB)
	if err != nil {
		return nil, fmt.Errorf("parse disk total failed: %w", err)
	}

	//返回的数据
	return &model.CmdInfo{
		CpuCore:   cpuCores,
		MemTotal:  memGBStr,
		DiskTotal: diskGBStr,
	}, nil
}

// 执行cmd命令
func runSSHCommand(conn *ssh.Client, cmd string) (string, error) {
	session, err := conn.NewSession()
	if err != nil {
		return "", err
	}
	defer func(session *ssh.Session) {
		err := session.Close()
		if err != nil {
			return
		}
	}(session)

	out, err := session.CombinedOutput(cmd)
	if err != nil {
		return "", fmt.Errorf("%w: %s", err, strings.TrimSpace(string(out)))
	}
	return string(out), nil
}
