package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"think-sso/api/v1/common"
	"think-sso/internal/model"
)

type UserListReq struct {
	g.Meta `path:"/user/list" tags:"用户列表" method:"get" summary:"用户列表"`
	common.PageReq
	common.Author
}

type UserListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.User `json:"list"`
	common.ListRes
}
