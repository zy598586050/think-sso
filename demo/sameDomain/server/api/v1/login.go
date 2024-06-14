package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model"
)

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
