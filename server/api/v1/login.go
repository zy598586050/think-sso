package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type EmailLoginReq struct {
	g.Meta   `path:"/login/email" tags:"Login" method:"post" summary:"登录"`
	Email    string `p:"email" v:"required#邮箱不能为空"`
	Password string `p:"password" v:"required#密码不能为空"`
}

type PhoneLoginReq struct {
	g.Meta   `path:"/login/phone" tags:"Login" method:"post" summary:"登录"`
	Email    string `p:"email" v:"required#邮箱不能为空"`
	Password string `p:"password" v:"required#密码不能为空"`
}

type LoginRes struct {
	g.Meta `mime:"application/json"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
}
