# gin_admin_api 接口
# 1. API接口
[http://127.0.0.1:5001/swagger/index.html](http://127.0.0.1:5001/swagger/index.html)
## 1.1. 登录接口
```go
http://localhost:5001/api/login
```

**接口参数:**

|    **参数** | 值 | 备注      |
| --- | --- |---------|
| username | wjs | 账号-test |
| password | wjs | 密码-test |
## 1.2. MYSQL数据库备份接口
```go
http://localhost:5001/api/backupMysql
数据库和ssh远程备份到本地的相关配置在config.yaml中
PS：该接口暂无参数，后续会实现自定义备份和相关配置，皆在UI界面操作
```