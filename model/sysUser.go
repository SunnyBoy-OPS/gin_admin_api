package model

// User 登录请求参数
type User struct {
	Username string `json:"username" gorm:"column:UserName"`
	Password string `json:"userpassword" gorm:"column:UserPassword"`
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
