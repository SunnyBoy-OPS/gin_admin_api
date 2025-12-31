// Package result 状态码和状态消息定义
package result

// Codes 定义状态
type Codes struct {
	Message map[uint]string
	Success uint
	Fail    uint
}

// ApiCode API Code状态码
var ApiCode = &Codes{
	Success: 200,
	Fail:    501,
}

// 状态信息初始化
func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.Success: "成功",
		ApiCode.Fail:    "失败",
	}
}

// GetMessage 供外部调用
func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
