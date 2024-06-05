package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"think-sso/api/v1/common"
	"think-sso/internal/model"
)

type EmailLoginReq struct {
	g.Meta   `path:"/login/email" tags:"邮箱登录" method:"post" summary:"邮箱登录"`
	Email    string `p:"email" v:"required#邮箱不能为空"`
	Password string `p:"password" v:"required#密码不能为空"`
}

type PhoneLoginReq struct {
	g.Meta   `path:"/login/phone" tags:"手机号登录" method:"post" summary:"手机号登录"`
	Email    string `p:"email" v:"required#邮箱不能为空"`
	Password string `p:"password" v:"required#密码不能为空"`
}

type LoginRes struct {
	g.Meta `mime:"application/json"`
	model.User
}

type LogoutReq struct {
	g.Meta `path:"/logout" tags:"退出登录" method:"post" summary:"退出登录"`
	common.Author
}

type LogoutRes struct {
	g.Meta `mime:"application/json"`
}
