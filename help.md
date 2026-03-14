# 导入yaml
go get gopkg.in/yaml.v2
go get gopkg.in/yaml.v3
# 导入日志logger
go get github.com/sirupsen/logrus

# 导入mysql 依赖
go get gorm.io/driver/mysql
# 导入gorm 依赖
go get gorm.io/gorm
# 远程连接MySQL 8.0，需要MySQL打开远程连接的权限
CREATE USER 'root'@'%' IDENTIFIED BY 'root';
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;

# 导入gin
go get -u github.com/gin-gonic/gin

# swaggo
//给你的 Gin 项目注册一个路由，用来访问 Swagger API 文档页面
go get github.com/swaggo/files
go get github.com/swaggo/gin-swagger
    # 下载swag
    go get -u github.com/swaggo/swag/cmd/swag
    # 安装一下swag
    go install github.com/swaggo/swag/cmd/swag@latest
//重新写入swag
swag init -g main.go -o docs 