// Package core mysql 初始化配置
// @author Flank
package core

import (
	"fmt"
	"gin_admin_api/config"
	"gin_admin_api/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

// MysqlInit mysql 连接
func MysqlInit() error {
	var err error
	var dbConfig = config.Config.Mysql
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Db,
		dbConfig.Charset)
	Db, err = gorm.Open(mysql.Open(url), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}
	if Db.Error != nil {
		return err
	}
	sqlDB, err := Db.DB()
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpen)
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdle)
	global.Log.Infof("[mysql]连接成功！")
	return nil
}
