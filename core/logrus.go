// 日志配置
// @author Flank

package core

import (
	"bytes"
	"fmt"
	"gin_admin_api/config"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

// 颜色
const (
	red    = 31
	green  = 32
	yellow = 33
	blue   = 34
	gray   = 36
)

type logFormatter struct{}

func (t *logFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 根据不同的level去展示不同的颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	log := config.Config.Logger

	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)

		//自定义输出格式
		fmt.Fprintf(
			b,
			"%s[%s] \x1b[%dm[%s]\x1b[0m %s %s: %s\n",
			log.Prefix,
			timestamp,
			levelColor,
			entry.Level.String(),
			funcVal,
			fileVal,
			entry.Message,
		)
	} else {
		//自定义输出格式
		fmt.Fprintf(
			b,
			"%s[%s] \x1b[%dm[%s]\x1b[0m %s\n",
			log.Prefix,
			timestamp,
			levelColor,
			entry.Level.String(),
			entry.Message,
		)
	}
	return b.Bytes(), nil
}

// InitLogger 初始化
func InitLogger() *logrus.Logger {
	mlog := logrus.New()                                // 新建一个实例
	mlog.SetOutput(os.Stdout)                           // 设置输出类型
	mlog.SetReportCaller(config.Config.Logger.ShowLine) // 开启返回的函数名和行号
	mlog.SetFormatter(&logFormatter{})                  // 设置自定义的Formatter
	level, err := logrus.ParseLevel(config.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mlog.SetLevel(level) //设置最低的Level
	InitDefaultLogger()  // 不注释即启用全局log
	return mlog
}

// InitDefaultLogger 全局log
func InitDefaultLogger() {
	logrus.SetOutput(os.Stdout)                           //设置输出类型
	logrus.SetReportCaller(config.Config.Logger.ShowLine) // 开启返回的函数名和行号
	logrus.SetFormatter(&logFormatter{})                  // 设置自定义的Formatter
	level, err := logrus.ParseLevel(config.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level) //设置最低的Level
}
