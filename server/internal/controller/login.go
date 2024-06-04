package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"think-sso/api/v1"
	"think-sso/internal/service"
	"time"
)

var (
	Login = cLogin{}
)

type cLogin struct{}

func (c *cLogin) EmailLogin(ctx context.Context, req *v1.EmailLoginReq) (res *v1.LoginRes, err error) {
	user, err := service.User().GetUserByEmailPassword(ctx, req)
	if user != nil {
		res = &v1.LoginRes{
			Id:    user.Id,
			Name:  user.Name,
			Phone: user.Phone,
			Email: user.Email,
		}
		token, err := service.Token().CreateToken(ctx, res)
		if err != nil {
			return nil, err
		}
		g.RequestFromCtx(ctx).Cookie.SetCookie("token", token, "/", "", g.Cfg().MustGet(ctx, "jwt.exp").Duration()*time.Minute)
	} else {
		err = gerror.New("密码错误")
	}
	return
}
