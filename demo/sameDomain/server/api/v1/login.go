package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model"
)

type CodeLoginReq struct {
	g.Meta `path:"/login/code" tags:"授权登录" method:"post" summary:"授权登录"`
	Code   string `p:"code" v:"required#code不能为空"`
}

type LoginRes struct {
	g.Meta `mime:"application/json"`
}

type UserInfoReq struct {
	g.Meta `path:"/user/info" tags:"用户信息" method:"post" summary:"用户信息"`
}

type UserInfoRes struct {
	g.Meta `mime:"application/json"`
	model.User
}

type LogoutReq struct {
	g.Meta `path:"/logout" tags:"退出登录" method:"post" summary:"退出登录"`
}

type LogoutRes struct {
	g.Meta `mime:"application/json"`
}
