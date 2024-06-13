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

// List 用户列表
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

// Info 用户信息
func (c *cUser) Info(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	user, err := service.Token().TokenToUser(ctx)
	if err != nil {
		return
	}
	res = &v1.UserInfoRes{
		User: *user,
	}
	return
}
