package model

// LoginReq 从请求体中获取账户密码
type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// User 登录请求参数
type User struct {
	Id       int    `json:"id" gorm:"column:Id"`
	Username string `json:"username" gorm:"column:UserName"`
	Password string `json:"userpassword" gorm:"column:UserPassword"`
	//Token    string `json:"token" binding:"required"`		//请求体
}

// TableName 重新映射表名，默认Gorm会查询users表
func (User) TableName() string {
	return "User"
}

// UserInfo 用户结构
type UserInfo struct {
	ID         int
	Username   string
	Password   string
	NickName   string
	Department string
	DeptName   string
}
