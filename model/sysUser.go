package model

// User 用户表结构体
type User struct {
	Username string `json:"username" gorm:"column:UserName"`
	Password string `json:"userpassword" gorm:"column:UserPassword"`
}

// TableName 重新映射表名，默认Gorm会查询users表
func (User) TableName() string {
	return "User"
}
