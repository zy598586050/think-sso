package controller

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"think-sso/api/v1"
	"think-sso/internal/model"
	"think-sso/internal/service"
	"think-sso/utility"
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

// Code 生成Code
func (c *cLogin) Code(ctx context.Context, req *v1.CodeReq) (res *v1.CodeRes, err error) {
	token, _ := utility.GetAuthorization(g.RequestFromCtx(ctx))
	prefix := g.Cfg().MustGet(ctx, "jwt.prefix").String()
	salt := g.Cfg().MustGet(ctx, "jwt.salt").String()
	// 加盐混淆加密
	code, _ := gmd5.EncryptString(token + salt)
	ttl, err := g.Redis().TTL(ctx, prefix+token)
	if err != nil {
		return
	}
	g.Redis().SetEX(ctx, prefix+code, token, ttl)
	res = &v1.CodeRes{
		Code: code,
	}
	return
}

func (c *cLogin) CodeLogin(ctx context.Context, req *v1.CodeLoginReq) (res *v1.LoginRes, err error) {
	prefix := g.Cfg().MustGet(ctx, "jwt.prefix").String()
	token, err := g.Redis().Get(ctx, prefix+req.Code)
	if err != nil {
		return
	}
	if token.String() == "" {
		err = gerror.New("code无效")
		return
	}
	ttl, _ := g.Redis().TTL(ctx, prefix+req.Code)
	userJson, _ := g.Redis().Get(ctx, prefix+token.String())
	user, err := gjson.DecodeToJson(userJson.String())
	if err != nil {
		return
	}
	res = &v1.LoginRes{
		User: model.User{
			Id:         user.Get("id").Int(),
			Name:       user.Get("name").String(),
			Avatar:     user.Get("avatar").String(),
			Phone:      user.Get("phone").String(),
			Email:      user.Get("email").String(),
			AppIds:     user.Get("appIds").String(),
			CreateTime: user.Get("createTime").GTime(),
			UpdateTime: user.Get("updateTime").GTime(),
		},
	}
	g.Redis().Del(ctx, prefix+req.Code)
	g.RequestFromCtx(ctx).Cookie.SetCookie("think-sso-token", token.String(), "", "/", time.Duration(ttl)*time.Second)
	return
}

func (c *cLogin) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	token, _ := utility.GetAuthorization(g.RequestFromCtx(ctx))
	err = service.Token().RemoveToken(ctx, token)
	return
}

func (c *cLogin) CheckAuth(ctx context.Context, req *v1.CheckAuthReq) (res *v1.CheckAuthRes, err error) {
	return
}
