package controller

import (
	"context"

	"think-sso/api/v1"
)

var (
	Login = cLogin{}
)

type cLogin struct{}

func (c *cLogin) Hello(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	return
}
