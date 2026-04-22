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
## 1.3. 主机管理接口
```go
http://localhost:5001/api/connect-test-host
PS：测试主机的连通性的接口
```
| **参数**   | 值    | 备注    |
|----------|------|-------|
| Host     | 主机ip | ssh连接 |
| Port     | 主机端口 | ssh连接 |
| Username | 账号   | ssh连接 |
| Password | 密码   | ssh连接 |
```go
http://localhost:5001/api/total-metric
PS：添加主机就自动采集主机的部分硬件信息
```
| **返回参数** | 值 | 备注  |
|----------|---|-----|
| CpuCore   |cpu核|ssh连接|
| MemTotal  |总内存|ssh连接|
| DiskTotal |总磁盘|ssh连接|

