package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta `path:"/login" tags:"Login" method:"post" summary:"登录"`
}

type LoginRes struct {
	g.Meta `mime:"application/json"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
}
