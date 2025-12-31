// 读取配置文件
// @author Flank

package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// config 总配置文件
type config struct {
	System system `json:"system"`
	Logger logger `json:"logger"`
	Mysql  mysql  `json:"mysql"`
}

// system系统配置
type system struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Env  string `yaml:"env"`
}

// logger 日志
type logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show_line"`
	LogInConsole bool   `yaml:"log_in_console"`
}

// 数据库MYSQL配置
type mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Loglevel string `yaml:"loglevel"`
	Charset  string `yaml:"charset"`
	MaxIdle  int    `yaml:"max_idle"`
	MaxOpen  int    `yaml:"max_open"`
}

var Config *config

// init 初始化配置
func init() {
	yamlFile, err := os.ReadFile("./config.yaml")
	if err != nil {
		return
	}

	// 初始化指针
	//Config = config{}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		return
	}
}
