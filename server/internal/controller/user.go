package controller

import (
	"context"
	v1 "think-sso/api/v1"
	"think-sso/api/v1/common"
	"think-sso/internal/service"
)

var (
	User = cUser{}
)

type cUser struct{}

func (c *cUser) List(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	total, userList, err := service.User().GetUserList(ctx, req)
	res = &v1.UserListRes{
		List: userList,
		ListRes: common.ListRes{
			Total: total,
		},
	}
	return
}
