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

type CodeLoginReq struct {
	g.Meta `path:"/login/code" tags:"授权登录" method:"post" summary:"授权登录"`
	Code   string `json:"code"`
}

type CodeReq struct {
	g.Meta `path:"/code" tags:"获取code" method:"post" summary:"获取code"`
	common.Author
}

type CodeRes struct {
	g.Meta `mime:"application/json"`
	Code   string `json:"code"`
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

type CheckAuthReq struct {
	g.Meta    `path:"/check/auth" tags:"验证权限" method:"post" summary:"验证权限"`
	AppId     string `p:"appId" v:"required#appId不能为空"`
	AppSecret string `p:"AppSecret" v:"required#AppSecret不能为空"`
	common.Author
}

type CheckAuthRes struct {
	g.Meta  `mime:"application/json"`
	IsLogin bool `json:"is_login"`
}
