package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"think-sso/api/v1"
	"think-sso/internal/model"
	"think-sso/internal/service"
	"think-sso/utility"
)

var (
	Login = cLogin{}
)

type cLogin struct{}

func (c *cLogin) EmailLogin(ctx context.Context, req *v1.EmailLoginReq) (res *v1.LoginRes, err error) {
	user, err := service.User().GetUserByEmailPassword(ctx, req)
	if user != nil {
		res = &v1.LoginRes{
			User: model.User{
				Id:         user.Id,
				Name:       user.Name,
				Avatar:     user.Avatar,
				Phone:      user.Phone,
				Email:      user.Email,
				AppIds:     user.AppIds,
				CreateTime: user.CreateTime,
				UpdateTime: user.UpdateTime,
			},
		}
		err = service.Token().CreateToken(ctx, &res.User)
	} else {
		err = gerror.New("密码错误")
	}
	return
}

func (c *cLogin) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	token, _ := utility.GetAuthorization(g.RequestFromCtx(ctx))
	err = service.Token().RemoveToken(ctx, token)
	return
}
