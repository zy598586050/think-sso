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
	model.User
}

type TestReq struct {
	g.Meta `path:"/test" tags:"测试" method:"get" summary:"测试"`
}

type TestRes struct {
	g.Meta `mime:"application/json"`
}
